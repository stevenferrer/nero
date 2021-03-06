package nero

import (
	"fmt"
)

// ErrRequiredField is a required field error
type ErrRequiredField struct {
	field string
}

// NewErrRequiredField returns an ErrFieldRequired error
func NewErrRequiredField(field string) *ErrRequiredField {
	return &ErrRequiredField{field: field}
}

func (e *ErrRequiredField) Error() string {
	return fmt.Sprintf("%s field is required", e.field)
}
