package email

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type VerificationMailer interface {
	SendMail(to string, view string, props map[string]string) error
}

type Gmail struct {
	SenderName string
	Dialer     *gomail.Dialer
}

func NewMailer(viper *viper.Viper) VerificationMailer {
	mailer := gomail.NewDialer(
		viper.GetString("MAILER_HOST"),
		viper.GetInt("MAILER_PORT"),
		viper.GetString("MAILER_EMAIL"),
		viper.GetString("MAILER_PASSWORD"),
	)

	sender := fmt.Sprintf("%s <%s>", viper.GetString("APP_NAME"), viper.GetString("MAILER_EMAIL"))

	return &Gmail{sender, mailer}
}

func (g *Gmail) SendMail(to string, view string, props map[string]string) error {
	tmpl, err := template.ParseFiles(view)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, props); err != nil {
		return err
	}

	email := gomail.NewMessage()
	email.SetHeader("From", g.SenderName)
	email.SetHeader("To", to)
	email.SetBody("text/html", buff.String())

	if err := g.Dialer.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
