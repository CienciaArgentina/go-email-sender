package commons

import (
	"errors"
	"fmt"
	"github.com/CienciaArgentina/go-email-sender/defines"
	"os"
	"reflect"
)

type ITemplateHelper interface {
	GetTemplateByName(template string, data interface{}) (*TemplateInfo, error)
	CheckIfTemplateFileExist(templateFile *string) bool
	CreateBodyFromInterface(template TemplateInfo, data interface{}) (*TemplateInfo, error)
}

type TemplateHelper struct {
}

func (t *TemplateHelper) GetTemplateByName(template string, data interface{}) (*TemplateInfo, error) {
	if IsNilOrEmpty(template) {
		return nil, errors.New(defines.TemplateNameCantBeEmpty)
	}

	templateToBeSent := TemplateMap[template]
	if templateToBeSent == (TemplateInfo{}) {
		return nil, errors.New(defines.TemplateDoesNotExist)
	}

	templateFileExists := t.CheckIfTemplateFileExist(templateToBeSent.Filename)
	if !templateFileExists {
		return nil, errors.New(defines.TemplateFileDoesNotExist)
	}

	mappedTemplate, err := t.CreateBodyForTemplate(templateToBeSent, data)
	if err != nil {
		return nil, err
	}

	return mappedTemplate, nil
}

func (t *TemplateHelper) CheckIfTemplateFileExist(templateFile string) bool {
	path := fmt.Sprintf("./templates/%s", templateFile)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (t *TemplateHelper) CreateBodyForTemplate(template TemplateInfo, data interface{}) (*TemplateInfo, error) {
	var err error
	switch template.Type {
	case defines.ConfirmEmail:
		template.Entity = &ConfirmationMailBody{TokenizedUrl: data.(string)}
	case defines.ForgotUsername:
		template.Entity = &ForgotUsernameBody{Username: data.(string)}
	}

	return &template, err
}

func (t *TemplateHelper) CreateBodyFromInterface(templateEntity interface{}, data interface{}) (interface{}, error) {
	if templateEntity == nil {
		return &templateEntity, nil
	}

	// Get the type and names of the properties of the templateEntity struct
	templateElems := reflect.ValueOf(templateEntity).Elem()
	templateTypes := templateElems.Type()
	templateColumns := templateElems.NumField()

	// Get the type and names of the properties of the data struct
	dataElems := reflect.ValueOf(data)
	dataTypes := dataElems.Type()
	if dataTypes.Kind() == reflect.Ptr && dataTypes.Elem().Kind() == reflect.Struct {
		dataTypes = dataTypes.Elem()
	} else {
		return nil, errors.New(defines.TemplateBodyMismatch)
	}
	dataColumns := dataTypes.NumField()

	// Check if the type is the same
	if templateElems.Kind() != dataElems.Kind() {
		return nil, errors.New(defines.TemplateBodyMismatch)
	}

	// Check if the property count is the same
	if templateColumns != dataColumns {
		return nil, errors.New(defines.TemplateBodyFieldsCountError)
	}

	// Iterate trough every field
	for n := 0; n < templateColumns; n++ {
		// Get the current field
		templateField := templateTypes.Field(n)
		currentTemplateField := templateElems.Field(n)

		// Find the property match
		for m := 0; m < dataColumns; m++ {
			dataField := dataTypes.Field(m)
			dataFieldElem := dataElems.Field(m).Interface()
			//currentDataField := dataElems.Field(m)
			if templateField.Name == dataField.Name {

				// Check if property is setteable
				if currentTemplateField.CanSet() {
					switch t := currentTemplateField.Interface().(type) {
					case string:
						currentTemplateField.SetString(dataFieldElem.(string))
						break
					case int64:
						currentTemplateField.SetInt(dataFieldElem.(int64))
						break
					case uint64:
						currentTemplateField.SetUint(dataFieldElem.(uint64))
						break
					case int32:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(int32)))
						break
					case uint32:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(uint32)))
						break
					case int16:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(int16)))
						break
					case uint16:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(uint16)))
						break
					case int8:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(int8)))
						break
					case uint8:
						currentTemplateField.Set(reflect.ValueOf(dataFieldElem.(uint8)))
						break
					case bool:
						currentTemplateField.SetBool(dataFieldElem.(bool))
						break
					default:
						return nil, errors.New(fmt.Sprintf("El tipo %s no pudo ser parseado. Deberíamos agregarlo al listado", t))
					}
				} else {
					return nil, errors.New(defines.CannotSetField)
				}
			}
		}
	}

	return templateEntity, nil
}
