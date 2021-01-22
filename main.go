package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/kulwindermatharu/go-smtp"
)

func main() {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "Test", "Test")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"send@send.com"}
	msg := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount !\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.live.com:587", auth, "mail.live.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
