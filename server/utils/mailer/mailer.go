package mailer

import (
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/mbaraa/apollo-music/config/env"
)

var (
	templates = template.Must(template.ParseGlob("./utils/mailer/*.html"))
)

func SendOTP(otp, to string) error {
	buf := new(strings.Builder)
	err := templates.ExecuteTemplate(buf, "otp", map[string]string{
		"OTP": otp,
	})
	if err != nil {
		return err
	}

	return sendEmail("Email verification", buf.String(), to)
}

func sendEmail(subject, content, to string) error {
	receiver := []string{to}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	_subject := "Subject: " + subject
	_to := "To: " + to
	_from := "From: Apollo Music"
	body := []byte(fmt.Sprintf("%s\n%s\n%s\n%s\n%s", _from, _to, _subject, mime, content))

	addr := env.MailerHost() + ":" + env.MailerPort()
	auth := smtp.PlainAuth("", env.MailerFrom(), env.MailerPassword(), env.MailerHost())

	return smtp.SendMail(addr, auth, env.MailerFrom(), receiver, body)
}
