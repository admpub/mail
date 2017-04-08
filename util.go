package mail

import (
	"net/mail"
)

func MailAddr(name string, address string) *mail.Address {
	return &mail.Address{
		Name:    name,
		Address: address,
	}
}

//SendMail 发送电邮
func SendMail(subject string, content string, receiver, sender string,
	attachments []string, smtpConfig *SMTPConfig, attachmentDir string) error {
	c := NewSMTPClient(smtpConfig)
	m := NewMail()
	m.AddTo(receiver) //receiver e.g. "Barry Gibbs <bg@example.com>"
	m.AddFrom(sender)
	m.AddSubject(subject)
	//m.AddText("Some text :)")
	m.AddHTML(content)
	m.BaseDir = attachmentDir
	if len(attachments) > 0 {
		for _, v := range attachments {
			m.AddAttachment(v)
		}
	}
	return c.Send(m)
}
