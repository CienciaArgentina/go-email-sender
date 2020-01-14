package main

import (
	"github.com/CienciaArgentina/go-email-sender/emailsender"
)

func main() {
	emailSenderController := emailsender.NewController()

	emailsender.InitRouter(emailSenderController)
}
