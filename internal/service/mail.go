package service

import (
	"github.com/dchest/uniuri"
	"github.com/resendlabs/resend-go"
	"go.uber.org/fx"
)

type Mail struct {
	fx.In

	ResendClient *resend.Client
}

type SendEmailRequest struct {
	To      []string
	Subject string
	Html    string
	Text    string
}

func (s Mail) Send(r *SendEmailRequest) (string, error) {
	resp, err := s.ResendClient.Emails.Send(&resend.SendEmailRequest{
		From:    "Penguin Statistics RogueStats <noreply+roguestats@penguin-stats.io>",
		To:      r.To,
		Subject: r.Subject,
		Html:    r.Html,
		Text:    r.Text,
		ReplyTo: "Penguin Statistics Support <support@penguin-stats.io>",
		Headers: map[string]string{
			"X-Entity-Ref-ID": uniuri.NewLen(16),
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
