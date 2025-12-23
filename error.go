package httplusplus

type FrameworkError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (err FrameworkError) Error() string {
	return err.Message
}

func SendError(status int, msg string) error {
	return FrameworkError{Status: status, Message: msg}
}
