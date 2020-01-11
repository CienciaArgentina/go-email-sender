package commons

type BaseResponse struct {
	Code    int
	Data    interface{}
	Error   string
	Message string
}

func NewBaseResponse(code int, data interface{}, err error, message string) *BaseResponse {
	response := BaseResponse{}
	response.Code = code
	response.Data = data
	if err != nil {
		response.Error = err.Error()
	}
	response.Message = message
	return &response
}
