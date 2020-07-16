package api


//https://blog.golang.org/error-handling-and-go
type BaseError struct {
	Error   error
	Message string
	Code    int
}
