package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	Host     string
	Port     int
	Username string
	dialer   *gomail.Dialer
	sender   *gomail.SendCloser
}

func GetTemplate(path string, body *bytes.Buffer, monthName, email string) {
	b := body
	t, err := template.ParseFiles(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(b, struct{ Month, Email string }{Month: monthName, Email: email})
}

func (e *EmailService) SetNewDialer() error {
	pass, err := goDotEnvVariable("PASSWORD")
	if err != nil {
		fmt.Println("Couldn't find the password")
		return err
	}

	e.dialer = gomail.NewDialer(e.Host, e.Port, e.Username, pass)
	return nil
}

func (e *EmailService) Connect() error {
	var s gomail.SendCloser
	var err error
	if s, err = e.dialer.Dial(); err != nil {
		return err
	}

	e.sender = &s
	return nil
}

func (e *EmailService) SendReceipt(email string, period, attachPath string, body *bytes.Buffer) error {
	var s gomail.SendCloser = *e.sender

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("soporte-administrativo@elmolio.net", "El Molio Soporte")},
		"To":      {email},
		"Subject": {"Recibo de mantenimiento " + period},
	})

	m.SetBody("text/html", body.String())
	m.Attach(attachPath)

	if err := gomail.Send(s, m); err != nil {
		return err
	}

	return nil
}

func (e *EmailService) Desconnect() error {
	s := *e.sender

	if err := s.Close(); err != nil {
		return err
	}

	return nil
}

func goDotEnvVariable(key string) (string, error) {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	return os.Getenv(key), nil
}

//
