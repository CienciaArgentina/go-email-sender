package commons

import (
	"github.com/CienciaArgentina/go-email-sender/defines"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckIfTemplateFileExistShouldReturnErrorWhenFileDoesNotExist(t *testing.T) {
	// Given
	fileName := "idonotexist.html"
	helper := TemplateHelper{}

	// When
	iHopeIDontExist := helper.CheckIfTemplateFileExist(fileName)

	// Then
	require.Equal(t, false, iHopeIDontExist)
}

func TestCreateBodyFromInterface(t *testing.T) {
	// Given
	expectedBody := ConfirmationMailBody{TokenizedUrl: "asd"}
	dto := NewDTO([]string{""}, expectedBody, defines.ConfirmEmail)
	helper := TemplateHelper{}

	// When
	template, _ := helper.GetTemplateByName(defines.ConfirmEmail, dto.Data)

	// Then
	require.Equal(t, &expectedBody, template.Entity)
}
