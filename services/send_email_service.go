package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/tokoku_backend/config"
)

type SmtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

var (
	err error
	doc bytes.Buffer
)

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`

func SendEmail() {
	auth := smtp.PlainAuth("",
		config.Config.Email.Username,
		config.Config.Email.Password,
		config.Config.Email.EmailServer,
	)

	context := &SmtpTemplateData{
		"SmtpEmailSender",
		"nikenprm@yahoo.com",
		"This is the email subject line",
		"Hello, this is a test email body.",
	}

	t := template.New("emailTemplate")
	t, err = t.Parse(emailTemplate)
	if err != nil {
		fmt.Print("error trying to parse mail template")
	}

	err = t.Execute(&doc, context)
	if err != nil {
		fmt.Print("error trying to execute mail template")
	}

	err = smtp.SendMail(config.Config.Email.EmailServer+":"+config.Config.Email.Port,
		auth,
		config.Config.Email.Username,
		[]string{"nikenprm@yahoo.com"},
		doc.Bytes(),
	)
	if err != nil {
		fmt.Print("ERROR: attempting to send a mail ", err)
	}

}
