package email

import "bytes"

type Email struct {
	From    string
	To      []string
	Subject string
	Body    bytes.Buffer
}
