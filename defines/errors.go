package defines

const (
	EmailAuthIsEmpty             = "Email, password o smtp faltante (configurame en las variables de entorno). Paniqueando la aplicación"
	DataCantBeNil                = "Data no puede ser nil"
	TemplateNameCantBeEmpty      = "El templatecommons no puede estar vacío"
	TemplateDoesNotExist         = "El templatecommons especificado no existe"
	TemplateFileDoesNotExist     = "El archivo de templatecommons que se quería usar no existe"
	TemplateBodyMismatch         = "La información del body recibido no matchea con el tipo esperado (ej. se envió una persona y se esperaba un auto)"
	TemplateBodyFieldsCountError = "La información del body podría ser del tipo esperado pero la cantidad de campos es distinta (quizás te olvidaste de cargarlo en template map?)"
	CannotSetField               = "El campo no es setteable. Por favor revisar si hay una variable addresseable"
)
