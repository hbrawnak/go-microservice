package main

import (
	"bytes"
	"html/template"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        interface{}
	DataMap     map[string]interface{}
}

func (m *Mail) SendSMTPMessage(message Message) error {
	if message.From == "" {
		message.From = m.FromAddress
	}

	if message.FromName == "" {
		message.FromName = m.FromName
	}

	data := map[string]interface{}{
		"message": message.Data,
	}

	message.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(message)
	if err != nil {
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(message)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClent, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(message.From).AddTo(message.To).SetSubject(message.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(message.Attachments) > 0 {
		for _, attachment := range message.Attachments {
			email.AddAttachment(attachment)
		}
	}

	if err := email.Send(smtpClent); err != nil {
		return err
	}

	return nil
}

func (m *Mail) buildHTMLMessage(message Message) (string, error) {
	templateToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err = t.ExecuteTemplate(&buffer, "body", message.DataMap); err != nil {
		return "", err
	}

	formattedMessage := buffer.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) inlineCSS(content string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(content, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *Mail) buildPlainTextMessage(message Message) (string, error) {
	templateToRender := "./templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err = t.ExecuteTemplate(&buffer, "body", message.DataMap); err != nil {
		return "", err
	}

	plainMessage := buffer.String()

	return plainMessage, nil
}

func (m *Mail) getEncryption(encryption string) mail.Encryption {
	switch encryption {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
