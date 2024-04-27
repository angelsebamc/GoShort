package http_status

type StatusCode int

type HTTPStatus struct {
	Code    StatusCode
	Message string
}

const (
	StatusOK           StatusCode = 200
	StatusCreated      StatusCode = 201
	StatusBadRequest   StatusCode = 400
	StatusUnauthorized StatusCode = 401
	StatusNotFound     StatusCode = 404
	StatusConflict     StatusCode = 409
	StatusInternal     StatusCode = 500
)

func NewHTTPStatus(statusCode StatusCode, message string) *HTTPStatus {
	return &HTTPStatus{Code: statusCode, Message: message}
}
