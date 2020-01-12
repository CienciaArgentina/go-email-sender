package emailsender

import (
	"errors"
	"github.com/CienciaArgentina/email-sender/commons"
	"github.com/CienciaArgentina/email-sender/defines"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type TemplateHelperMock struct {
	mock.Mock
}

func (t *TemplateHelperMock) GetTemplateByName(template string, data interface{}) (*commons.TemplateInfo, error) {
	args := t.Called(template, data)
	return args.Get(0).(*commons.TemplateInfo), args.Error(1)
}

func (t *TemplateHelperMock) CheckIfTemplateFileExist(templateFile string) bool {
	return false
}

func (t *TemplateHelperMock) CreateBodyFromInterface(templateEntity interface{}, data interface{}) (interface{}, error) {
	return nil, nil
}

func (t *TemplateHelperMock) CreateBodyForTemplate(template commons.TemplateInfo, data interface{}) (*commons.TemplateInfo, error) {
	return nil, nil
}

func TestInvokeEmailSenderShouldReturnErrorWhenDTOIsNil(t *testing.T) {
	// Given
	var dto commons.DTO
	service := NewService()

	// When
	result := service.InvokeEmailSender(dto)

	// Then
	require.Equal(t, http.StatusBadRequest, result.Code)
}

func TestParseTemplateShouldReturnErrorWhenTemplateDoesNotExist(t *testing.T) {
	// Given
	helperMock := new(TemplateHelperMock)
	service := NewService(helperMock)
	dto := createDtoWithData()
	helperMock.On(defines.GetTemplateByName, dto.Template, dto.Data).Return(&commons.TemplateInfo{}, errors.New(defines.TemplateDoesNotExist))

	// When
	result := service.ParseTemplate(dto)

	// Then
	require.Equal(t, http.StatusBadRequest, result.Code)
}

func createDtoWithData() commons.DTO {
	return commons.DTO{
		To:       []string {"juan@carlos.com"},
		Data:     commons.ConfirmationMailBody{
			Name:  "Juan",
			Token: "T0K3N",
		},
		Template: defines.ConfirmEmail,
	}
}