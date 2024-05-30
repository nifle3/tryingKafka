package email

import (
	"bytes"

	"consumerSMTP/internal/entities"
)

type Templater interface {
	Render(data []entities.Currency) (bytes.Buffer, error)
}

type Builder struct {
	email    Email
	template Templater
}

func New(template Templater) *Builder {
	return &Builder{
		template: template,
	}
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

func (b *Builder) SetBody(currencies []entities.Currency) error {
	var err error

	b.email.Body, err = b.template.Render(currencies)
	return err
}

func (b *Builder) Build() Email {
	return b.email
}
