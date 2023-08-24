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
	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type Auth struct {
	fx.In

	Config      *appconfig.Config
	UserRepo    repo.User
	JWT         JWT
	Turnstile   Turnstile
	MailService Mail
}

func (s Auth) AuthByLoginInput(ctx context.Context, args model.LoginInput) (*model.User, error) {
	err := s.Turnstile.Verify(ctx, args.TurnstileResponse, "login")
	if err != nil {
		return nil, gqlerror.Errorf("captcha verification failed: invalid turnstile response")
	}

	user, err := s.UserRepo.GetUserByEmail(ctx, args.Email)
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

func (s Auth) AuthByToken(ctx context.Context, token string) (*model.User, error) {
	userId, expireAt, err := s.JWT.Verify(token)
	if err != nil {
		return nil, err
	}

	// auto renew token
	if time.Until(expireAt) < s.Config.JWTAutoRenewalTime {
		err = s.SetUserToken(ctx, &model.User{ID: userId})
		if err != nil {
			return nil, err
		}
	}

	return s.UserRepo.GetUserByID(ctx, userId)
}

func (s Auth) CreateUser(ctx context.Context, args model.CreateUserInput) (*model.User, error) {
	var randomBytes [32]byte
	_, err := rand.Read(randomBytes[:])
	if err != nil {
		return nil, err
	}
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes[:])

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomString), 12)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:       args.Name,
		Email:      &args.Email,
		Attributes: args.Attributes,
		Credential: string(hashedPassword),
	}

	err = s.UserRepo.CreateUser(ctx, user)
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
		To:      []string{*user.Email},
		Subject: "你的 RogueStats 登录信息已准备就绪",
		Html:    rendered.HTML,
		Text:    rendered.Text,
	})
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s Auth) CurrentUser(ctx context.Context) (*model.User, error) {
	user := appcontext.CurrentUser(ctx)
	if user != nil {
		return user, nil
	}

	return nil, errors.New("you are not logged in")
}

func (s Auth) SetUserToken(ctx context.Context, user *model.User) error {
	fiberCtx := ctx.Value("fiberCtx").(*fiber.Ctx)

	token, err := s.JWT.Sign(user.ID)
	if err != nil {
		return err
	}

	fiberCtx.Set("X-Penguin-RogueStats-Set-Token", token)

	return nil
}
