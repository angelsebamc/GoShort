package json_response

import "goshort/utils/http_status"

type JSONResponse struct {
	Code    http_status.StatusCode `json:"code"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
}

func New(code http_status.StatusCode, message string, data interface{}) *JSONResponse {
	return &JSONResponse{Code: code, Message: message, Data: data}
}
