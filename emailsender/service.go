package emailsender

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/CienciaArgentina/go-email-sender/commons"
	"github.com/CienciaArgentina/go-email-sender/defines"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
)

type IEmailSenderService interface {
	GetAuth() smtp.Auth
	InvokeEmailSender(dto commons.DTO) *commons.BaseResponse
	ParseTemplate(dto commons.DTO) *commons.BaseResponse
	SendEmail(dto commons.DTO) *commons.BaseResponse
}

type EmailSenderService struct {
	TemplateHelper commons.TemplateHelper
	EmailFormat
}

type EmailFormat struct {
	TemplateInfo *commons.TemplateInfo
	Body         string
}

func (e *EmailSenderService) GetAuth() smtp.Auth {

	username := os.Getenv(defines.CienciaArgentinaEmail)
	password := os.Getenv(defines.CienciaArgentinaPassword)
	mailSmtp := os.Getenv(defines.CienciaArgentinaEmailSmtp)

	if commons.IsNilOrEmpty(username) || commons.IsNilOrEmpty(password) || commons.IsNilOrEmpty(mailSmtp) {
		panic(defines.EmailAuthIsEmpty)
	}
	auth := smtp.PlainAuth(
		defines.Identity,
		"noreply@cienciaargentina.com",
		password,
		"smtp.zoho.com",
	)

	return auth
}

func (e *EmailSenderService) InvokeEmailSender(dto commons.DTO) *commons.BaseResponse {
	if &dto == (commons.NewDTO(nil, nil, "")) {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, errors.New(defines.DataCantBeNil), defines.StringEmpty)
		return result
	}

	parseTemplateResult := e.ParseTemplate(dto)

	if parseTemplateResult.Error != "" {
		return parseTemplateResult
	}

	emailSendResult := e.SendEmail(dto)

	if emailSendResult != nil {
		return emailSendResult
	}

	return commons.NewBaseResponse(http.StatusOK, nil, nil, "")
}

func (e *EmailSenderService) ParseTemplate(dto commons.DTO) *commons.BaseResponse {
	var err error
	var result commons.BaseResponse
	e.TemplateInfo, err = e.TemplateHelper.GetTemplateByName(dto.Template, dto.Data)
	if err != nil {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty)
		return result
	}

	template, err := template.ParseFiles(fmt.Sprintf("./templates/%s", e.TemplateInfo.Filename))
	if err != nil {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty)
		return result
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, e.TemplateInfo.Entity); err != nil {
		return commons.NewBaseResponse(http.StatusInternalServerError, nil, err, defines.StringEmpty)
	}

	e.Body = buf.String()

	return &result
}

func (e *EmailSenderService) SendEmail(dto commons.DTO) *commons.BaseResponse {
	//// Setup headers
	//headers := make(map[string]string)
	//headers["From"] = "noreply@cienciaargentina.com"
	//headers["To"] = dto.To[0]
	//headers["Subject"] = e.TemplateInfo.Subject
	//
	////message := ""
	////for k,v := range headers {
	////	message += fmt.Sprintf("%s: %s\r\n", k, v)
	////}
	////message += "\r\n" + e.Body

	servername := "smtp.zoho.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("","noreply@cienciaargentina.com", os.Getenv(defines.CienciaArgentinaPassword), host)

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail("noreply@cienciaargentina.com"); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(dto.To[0]); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	formattedMsg := []byte(fmt.Sprintf("Subject: %s\n%s\n\n%s", e.TemplateInfo.Subject, defines.Mime, e.Body))

	_, err = w.Write(formattedMsg)
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}


	//if err := smtp.SendMail("smtp.zoho.com:587", e.GetAuth(), "noreply@cienciaargentina.com", dto.To, formattedMsg); err != nil {
	//	return commons.NewBaseResponse(http.StatusInternalServerError, nil, err, defines.StringEmpty)
	//}

	return nil
}

func NewService() *EmailSenderService {
	return &EmailSenderService{
		TemplateHelper: commons.TemplateHelper{},
	}
}
