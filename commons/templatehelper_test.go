package commons

import (
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
