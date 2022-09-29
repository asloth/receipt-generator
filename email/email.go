package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func GetTemplate(path string, body *bytes.Buffer, monthName string) {
	b := body
	t, err := template.ParseFiles(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(b, struct{ Month string }{Month: monthName})
}
func sendReceiptEmail(email string, clientName, templatePath, period, attachPath string) error {
	pass, err := goDotEnvVariable("PASSWORD")
	if err != nil {
		fmt.Println("Couldn't find the password")
		return err
	}

	var body bytes.Buffer
	GetTemplate(templatePath, &body, period)

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("soporte-administrativo@elmolio.net", "El Molio Soporte")},
		"To":      {email},
		"Subject": {"Recibo de mantenimiento " + period},
	})
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")

	m.SetBody("text/html", body.String())
	m.Attach(attachPath)

	d := gomail.NewDialer("smtp.gmail.com", 587, "soporte-administrativo@elmolio.net", pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func goDotEnvVariable(key string) (string, error) {

	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	return os.Getenv(key), nil
}

// sendReceiptEmail("sbenelramirez@gmail.com", "Sara", "./templates/maintenance.html", "Agosto-2022", "../GPR-RECIBOS-SEPTIEMBRE-2022/MANTENIMIENTO-SEPTIEMBRE-2022_DPTO-1910.pdf")
