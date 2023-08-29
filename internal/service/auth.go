package service

import (
	"context"
	"time"

	"github.com/dchest/uniuri"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vektah/gqlparser/v2/gqlerror"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"exusiai.dev/roguestats-backend/internal/appcontext"
	"exusiai.dev/roguestats-backend/internal/blob"
	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/ent/user"
	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/x/rediskey"
)

type Auth struct {
	fx.In

	Config      *appconfig.Config
	Ent         *ent.Client
	Redis       *redis.Client
	JWT         JWT
	Turnstile   Turnstile
	MailService Mail
}

func (s Auth) AuthByLoginInput(ctx context.Context, args model.LoginInput) (*ent.User, error) {
	err := s.Turnstile.Verify(ctx, args.TurnstileResponse, "login")
	if err != nil {
		return nil, err
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
	randomString := uniuri.NewLen(24)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomString), 12)
	if err != nil {
		return nil, err
	}

	client := ent.FromContext(ctx)

	user, err := client.User.Create().
		SetName(args.Name).
		SetEmail(args.Email).
		SetAttributes(args.Attributes).
		SetCredential(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	rendered, err := blob.RenderPair("password-generated", map[string]any{
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
	fiberCtx := appcontext.FiberCtx(ctx)

	token, err := s.JWT.Sign(user.ID)
	if err != nil {
		return err
	}

	fiberCtx.Set("X-Penguin-RogueStats-Set-Token", token)

	return nil
}

func (s Auth) RequestPasswordReset(ctx context.Context, input model.RequestPasswordResetInput) (bool, error) {
	err := s.Turnstile.Verify(ctx, input.TurnstileResponse, "reset-password")
	if err != nil {
		return false, err
	}
	user, err := s.Ent.User.Query().
		Where(user.Email(input.Email)).
		First(ctx)
	if err != nil {
		return false, err
	}

	resetToken := user.ID + "_" + uniuri.NewLen(32)

	err = s.Redis.Set(ctx, rediskey.ResetToken(resetToken), user.ID, s.Config.PasswordResetTokenTTL).Err()
	if err != nil {
		return false, err
	}

	rendered, err := blob.RenderPair("password-reset", map[string]any{
		"Username": user.Name,
		"Token":    resetToken,
		"TokenTTL": s.Config.PasswordResetTokenTTL.String(),
	})
	if err != nil {
		return false, err
	}

	_, err = s.MailService.Send(&SendEmailRequest{
		To:      []string{user.Email},
		Subject: "你的 RogueStats 登录信息重置请求已就绪",
		Html:    rendered.HTML,
		Text:    rendered.Text,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s Auth) ResetPassword(ctx context.Context, input model.ResetPasswordInput) (bool, error) {
	cmd := s.Redis.Get(ctx, rediskey.ResetToken(input.Token))
	if cmd.Err() != nil {
		return false, gqlerror.Errorf("invalid password reset token: the token is either expired, invalid or has been used")
	}

	user, err := s.Ent.User.Query().
		Where(user.ID(cmd.Val())).
		First(ctx)
	if err != nil {
		return false, err
	}

	err = passwordvalidator.Validate(input.Password, 60)
	if err != nil {
		return false, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return false, err
	}

	_, err = s.Ent.User.UpdateOneID(user.ID).
		SetCredential(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		return false, err
	}

	if s.Redis.Del(ctx, rediskey.ResetToken(input.Token)).Err() != nil {
		return false, err
	}

	return true, nil
}
