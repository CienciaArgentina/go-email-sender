package main

import (
	"github.com/CienciaArgentina/email-sender/emailsender"
)

func main() {
	emailSenderController := emailsender.NewController()

	emailsender.InitRouter(emailSenderController)
}
