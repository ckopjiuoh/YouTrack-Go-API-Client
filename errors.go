package ytapi

// errorString is a trivial implementation of error.
type NotFoundError struct {
	s string
}

func (e *NotFoundError) Error() string {
	return e.s
}

var NotFound = &NotFoundError{
	s: "Object not found",
}
