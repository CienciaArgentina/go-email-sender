package emailsender

import (
	"github.com/CienciaArgentina/email-sender/commons"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestInvokeEmailSenderShouldReturnErrorWhenDTOIsNil(t *testing.T) {
	// Given
	var dto commons.DTO
	service := NewService()

	// When
	result := service.InvokeEmailSender(dto)

	// Then
	require.Equal(t, http.StatusBadRequest, result.Code)
}
