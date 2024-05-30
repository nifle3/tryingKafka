package sender

import (
	"consumerSMTP/internal/email"
	"context"
	"net/smtp"

	"consumerSMTP/internal/entities"
)

type SMTP struct {
	auth smtp.Auth
	addr string
	from string
	to   []string
}

func New(auth smtp.Auth, addr, from string) *SMTP {
	return &SMTP{
		auth: auth,
		addr: addr,
		from: from,
	}
}

func (s *SMTP) Start(ctx context.Context, info <-chan entities.Message) {
	for {
		select {
		case msg := <-info:
			s.Send(ctx, msg)
		case <-ctx.Done():
			return
		}
	}
}

func (s *SMTP) Send(_ context.Context, msg entities.Message) {
	builder := email.New()
	builder.SetFrom(s.from)
	builder.SetSubject(msg.Id)
	builder.SetBody(msg.Result)

	for _, value := range s.to {
		builder.SetTo(value)

		mail := builder.Build()
		if err := smtp.SendMail(s.addr, s.auth, mail.From, mail.To, mail.Body.Bytes()); err != nil {
			continue
		}
	}
}
