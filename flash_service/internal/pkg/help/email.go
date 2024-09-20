package help

import (
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Params struct {
	From     string
	Password string
	To       string
	Message  string
	Code     string
	UserName string
}

func SendVerificationCode(params Params) error {
	htmlFile, err := os.ReadFile("internal/pkg/help/format.html")
	if err != nil {
		log.Println("Cannot read HTML file at path 'internal/pkg/help/format.html':", err)
		return err
	}
	temp, err := template.New("email").Parse(string(htmlFile))
	if err != nil {
		log.Println("Cannot parse HTML file:", err)
		return err
	}

	var Builder strings.Builder
	err = temp.Execute(&Builder, params)
	if err != nil {
		log.Println("Cannot execute HTML template:", err)
		return err
	}

	message := "From: " + params.From + "\n" +
		"To: " + params.To + "\n" +
		"Subject: Notification\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
		Builder.String()

	auth := smtp.PlainAuth("", params.From, params.Password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, params.From, []string{params.To}, []byte(message))
	if err != nil {
		log.Println("Could not send an email:", err)
		return err
	}

	return nil
}
