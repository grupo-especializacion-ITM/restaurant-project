package email

import (
	"log"
	"net/smtp"

	"kafka.consumer.go/src/config"
)

func SendEmail(conf *config.Conf, subject, body string) error {

	to := []string{conf.EmailTo}

	msg := "From: " + conf.EmailFrom + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", conf.EmailApiKey, conf.EmailSecretKey, "in-v3.mailjet.com")

	err := smtp.SendMail("in-v3.mailjet.com:587", auth, conf.EmailFrom, to, []byte(msg))
	if err != nil {
		log.Println("❌ Error al enviar el correo:", err)
		return err
	}
	log.Println("📧 Email enviado:", subject)
	return nil
}
