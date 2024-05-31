package helper

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"

	"github.com/joho/godotenv"
)

func SendEmail(email string, token string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "123200103@student.upnyk.ac.id")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verify Token")
	m.SetBody("text/html", "Hello, this is your token: "+token)

	d := gomail.NewDialer("smtp.gmail.com", 587, "123200103@student.upnyk.ac.id", os.Getenv("EMAIL_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	return nil
}
