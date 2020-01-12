package emailsender

import (
	"bytes"
	"errors"
	"github.com/CienciaArgentina/email-sender/commons"
	"github.com/CienciaArgentina/email-sender/defines"
	"html/template"
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

type EmailSenderService struct{
	TemplateHelper commons.ITemplateHelper
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
		username,
		password,
		mailSmtp,
	)

	return auth
}

func (e *EmailSenderService) InvokeEmailSender(dto commons.DTO) *commons.BaseResponse {
	if &dto == (commons.NewDTO(nil, nil, "")) {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, errors.New(defines.DataCantBeNil), defines.StringEmpty)
		return result
	}

	parseTemplateResult := e.ParseTemplate(dto)

	if parseTemplateResult.Code != http.StatusOK {
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
	e.TemplateInfo, err = e.TemplateHelper.GetTemplateByName(dto.Template, dto.Data)
	if err != nil {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty)
		return result
	}

	template, err := template.ParseFiles(e.TemplateInfo.Filename)
	if err != nil {
		result := commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty)
		return result
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, e.TemplateInfo.Entity); err != nil {
		return commons.NewBaseResponse(http.StatusInternalServerError, nil, err, defines.StringEmpty)
	}

	e.Body = buf.String()

	return commons.NewBaseResponse(http.StatusOK, nil, nil, defines.StringEmpty)
}

func (e *EmailSenderService) SendEmail(dto commons.DTO) *commons.BaseResponse {
	formattedMsg := []byte(e.TemplateInfo.Subject + defines.Mime + "\n" + e.Body)
	if err := smtp.SendMail(defines.CienciaArgentinaEmailSmtpPort, e.GetAuth(), os.Getenv(defines.CienciaArgentinaEmail), dto.To, formattedMsg); err != nil {
		return commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty)
	}

	return nil
}

func NewService(templateHelper ...commons.ITemplateHelper) *EmailSenderService {

	if len(templateHelper) == 0 {
		templateHelper = append(templateHelper, commons.NewHelper())
	}

	return &EmailSenderService{
		TemplateHelper: templateHelper[0],
	}
}
