package emailsender

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"github.com/CienciaArgentina/go-email-sender/commons"
	"github.com/CienciaArgentina/go-email-sender/defines"
	"html/template"
	"net"
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

func (e *EmailSenderService) InvokeEmailSender(dto commons.DTO) apierror.ApiError {
	if &dto == (commons.NewDTO(nil, nil, "")) {
		return apierror.NewBadRequestApiError(defines.DataCantBeNil)
	}

	parseTemplateResult := e.ParseTemplate(dto)

	if parseTemplateResult != nil {
		return parseTemplateResult
	}

	emailSendResult := e.SendEmail(dto)

	if emailSendResult != nil {
		return emailSendResult
	}

	return nil
}

func (e *EmailSenderService) ParseTemplate(dto commons.DTO) apierror.ApiError {
	var err error
	e.TemplateInfo, err = e.TemplateHelper.GetTemplateByName(dto.Template, dto.Data)
	if err != nil {
		clog.Error("GetTemplateByName error", "parse-template", err, nil)
		return apierror.NewBadRequestApiError(err.Error())
	}

	template, err := template.ParseFiles(fmt.Sprintf("./templates/%s", e.TemplateInfo.Filename))
	if err != nil {
		clog.Error("ParseFiles error", "parse-template", err, nil)
		return apierror.NewBadRequestApiError(err.Error())
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, e.TemplateInfo.Entity); err != nil {
		clog.Error("Execute template error", "parse-template", err, nil)
		return apierror.NewInternalServerApiError(err.Error(), err, "execute_template")
	}

	e.Body = buf.String()

	return nil
}

func (e *EmailSenderService) SendEmail(dto commons.DTO) apierror.ApiError {
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
		clog.Error("Dial error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Dial error", err, "email_error")
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		clog.Error("Error starting email client", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Error starting email client", err, "email_error")
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		clog.Error("Auth error", "send-email", err, nil)
	}

	// To && From
	if err = c.Mail("noreply@cienciaargentina.com"); err != nil {
		clog.Error("Mail error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Mail error", err, "email_error")
	}

	if err = c.Rcpt(dto.To[0]); err != nil {
		clog.Error("Receipt error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Receipt error", err, "email_error")
	}

	// Data
	w, err := c.Data()
	if err != nil {
		clog.Error("Data error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Data error", err, "email_error")
	}

	formattedMsg := []byte(fmt.Sprintf("Subject: %s\n%s\n\n%s", e.TemplateInfo.Subject, defines.Mime, e.Body))

	_, err = w.Write(formattedMsg)
	if err != nil {
		clog.Error("Write error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Write error", err, "email_error")
	}

	err = w.Close()
	if err != nil {
		clog.Error("Close error", "send-email", err, nil)
		return apierror.NewInternalServerApiError("Close error", err, "email_error")
	}
	return nil
}

func NewService() *EmailSenderService {
	return &EmailSenderService{
		TemplateHelper: commons.TemplateHelper{},
	}
}
