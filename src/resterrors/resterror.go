package resterrors

type RestError struct {
	Error      string
	Message    string
	HttpStatus int
}
