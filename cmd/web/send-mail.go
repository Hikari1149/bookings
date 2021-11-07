package main

import (
	"log"
	"time"

	"github.com/Hikari1149/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {

	go func() {

		for {

			msg := <-app.MailChan

			sendMsg(msg)

		}

	}()
}

func sendMsg(m models.MailData) {

	server := mail.NewSMTPClient()

	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subuject)
	// email.SetBody(mail.TextHTML, "Hello, <strong>HIkari</strong> !")
	email.SetBody(mail.TextHTML, m.Congtent)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}

}
