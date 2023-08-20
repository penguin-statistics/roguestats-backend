package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"exusiai.dev/roguestats-backend/internal/appcontext"
)

type Turnstile struct {
	Config *appconfig.Config

	httpClient *http.Client
}

func NewTurnstile(config *appconfig.Config) Turnstile {
	return Turnstile{
		Config: config,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s Turnstile) Verify(ctx context.Context, turnstileResponse, expectedAction string) error {
	if turnstileResponse == "" {
		return errors.New("missing turnstile response")
	}

	fiberCtx := appcontext.FiberCtx(ctx)

	body := url.Values{}
	body.Set("secret", s.Config.TurnstileSecret)
	body.Set("response", turnstileResponse)
	body.Set("remoteip", fiberCtx.IP())

	bodyReader := strings.NewReader(body.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://challenges.cloudflare.com/turnstile/v0/siteverify", bodyReader)
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	type turnstileVerificationResponse struct {
		Success bool   `json:"success"`
		Action  string `json:"action"`
	}

	var verificationResponse turnstileVerificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&verificationResponse); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}

	if !verificationResponse.Success {
		return errors.New("failed to verify captcha")
	}

	if expectedAction != "" && verificationResponse.Action != expectedAction {
		return errors.New("invalid captcha action")
	}

	return nil
}
