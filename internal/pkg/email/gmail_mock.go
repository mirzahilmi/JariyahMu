package email

import "github.com/stretchr/testify/mock"

type MailerMock struct {
	mock.Mock
}

func Foo() VerificationMailer {
	return &MailerMock{}
}

func (m *MailerMock) SendMail(to string, subject string, view string, props map[string]string) error {
	// TODO: Proper mock
	return nil
}
