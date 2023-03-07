package errors

const NotFoundMessage = "not found"

type NotFoundError struct {
	Message string
	Code    int
}

func (e *NotFoundError) Error() string {
	if e.Message == "" {
		return NotFoundMessage
	}
	return e.Message
}
