package utils

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func ErroFile() error {
	return &CustomError{Message: "Occurred when opening the file"}
}

func WriteFile() error {
	return &CustomError{Message: "Error writing the file"}
}
