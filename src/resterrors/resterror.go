package resterrors

type RestError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HttpStatus int    `json:"httpstatus"`
}
