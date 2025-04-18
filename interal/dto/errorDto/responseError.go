package errorDto

type ResponsError struct {
	Message string `json:"message"`
}

func NewResponsError(message string, err error) *ResponsError {
	return &ResponsError{message + ": " + err.Error()}
}
