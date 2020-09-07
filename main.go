package main

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"github.com/CienciaArgentina/go-email-sender/emailsender"
)

func main() {
	clog.SetLogLevel(clog.DebugLevel)
	emailSenderController := emailsender.NewController()
	emailsender.InitRouter(emailSenderController)
}
