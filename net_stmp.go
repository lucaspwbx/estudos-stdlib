package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	serverInfo := &smtp.ServerInfo{
		Name: "dokken",
		TLS:  true,
		Auth: []string{"foo", "bar"},
	}
	fmt.Println(serverInfo)

	//plain auth implements AuthInterface
	auth := smtp.PlainAuth("bla", "foo", "bar", "foobaz.com")
	fmt.Println(auth)

	//crammd5 auth
	crammd5 := smtp.CRAMMD5Auth("foo", "secret")
	fmt.Println(crammd5)

	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err) //no such host
	}
	err = c.Mail("sender@example.org")
	if err != nil {
		log.Fatal(err)
	}
	err := c.Rcpt("foo@bar.com")
	if err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}

	//send email
	auth3 := smtp.PlainAuth("", "foo@bar.com", "password", "mail.example.com")
	to := []string{"bla@bla.com"}
	msg := []byte("body")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
