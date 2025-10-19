package mailer

import (
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type SMTPClient struct {
	MailConfig
	Server *mail.SMTPServer
}

func (m *SMTPClient) InitServer() error {

	m.Server = mail.NewSMTPClient()
	m.Server.Host = m.Host
	m.Server.Port = m.Port
	m.Server.Username = m.Username
	m.Server.Password = m.Password
	m.Server.Encryption = m.getEncryption(m.Encryption)
	m.Server.KeepAlive = false
	m.Server.ConnectTimeout = 10 * time.Second
	m.Server.SendTimeout = 10 * time.Second
	return nil
}
func (m *SMTPClient) Send(msg Message, isSandbox bool) error {

	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		return err
	}

	client, err := m.Server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject)

	email.SetBody(mail.TextHTML, formattedMessage)
	email.AddAlternative(mail.TextPlain, plainMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}

	err = email.Send(client)
	if err != nil {
		return err
	}

	return nil

}
