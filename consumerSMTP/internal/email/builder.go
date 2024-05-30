package email

import (
	"consumerSMTP/internal/entities"
)

type Builder struct {
	email Email
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) SetFrom(from string) {
	b.email.From = from
}

func (b *Builder) SetTo(to string) {
	b.email.To = []string{to}
}

func (b *Builder) SetSubject(subject string) {
	b.email.Subject = subject
}

func (b *Builder) SetBody(currencies []entities.Currency) {
}

func (b *Builder) Build() Email {
	return b.email
}
