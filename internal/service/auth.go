package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"exusiai.dev/roguestats-backend/internal/appcontext"
	"exusiai.dev/roguestats-backend/internal/blob"
	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/ent/user"
	"exusiai.dev/roguestats-backend/internal/model"
)

type Auth struct {
	fx.In

	Config      *appconfig.Config
	Ent         *ent.Client
	JWT         JWT
	Turnstile   Turnstile
	MailService Mail
}

func (s Auth) AuthByLoginInput(ctx context.Context, args model.LoginInput) (*ent.User, error) {
	err := s.Turnstile.Verify(ctx, args.TurnstileResponse, "login")
	if err != nil {
		return nil, gqlerror.Errorf("captcha verification failed: invalid turnstile response")
	}

	user, err := s.Ent.User.Query().
		Where(user.Email(args.Email)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Credential), []byte(args.Password))
	if err != nil {
		return nil, err
	}

	err = s.SetUserToken(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s Auth) AuthByToken(ctx context.Context, token string) (*ent.User, error) {
	userId, expireAt, err := s.JWT.Verify(token)
	if err != nil {
		return nil, err
	}

	// auto renew token
	if time.Until(expireAt) < s.Config.JWTAutoRenewalTime {
		err = s.SetUserToken(ctx, &ent.User{ID: userId})
		if err != nil {
			return nil, err
		}
	}

	return s.Ent.User.Get(ctx, userId)
}

func (s Auth) CreateUser(ctx context.Context, args model.CreateUserInput) (*ent.User, error) {
	var randomBytes [16]byte
	_, err := rand.Read(randomBytes[:])
	if err != nil {
		return nil, err
	}
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes[:])

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomString), 12)
	if err != nil {
		return nil, err
	}

	user, err := s.Ent.User.Create().
		SetName(args.Name).
		SetEmail(args.Email).
		SetAttributes(args.Attributes).
		SetCredential(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	rendered, err := blob.RenderPair("password-generated", map[string]interface{}{
		"Username": user.Name,
		"Email":    user.Email,
		"Password": randomString,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.MailService.Send(&SendEmailRequest{
		To:      []string{user.Email},
		Subject: "你的 RogueStats 登录信息已准备就绪",
		Html:    rendered.HTML,
		Text:    rendered.Text,
	})
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s Auth) CurrentUser(ctx context.Context) (*ent.User, error) {
	user := appcontext.CurrentUser(ctx)
	if user != nil {
		return user, nil
	}

	return nil, errors.New("you are not logged in")
}

func (s Auth) SetUserToken(ctx context.Context, user *ent.User) error {
	fiberCtx := ctx.Value("fiberCtx").(*fiber.Ctx)

	token, err := s.JWT.Sign(user.ID)
	if err != nil {
		return err
	}

	fiberCtx.Set("X-Penguin-RogueStats-Set-Token", token)

	return nil
}
