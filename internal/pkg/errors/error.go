package errors

import (
	"fmt"
)

func ErrInvalidFieldFormat(customMsg string, err error) error {
	return fmt.Errorf("invalid field format: %s", customMsg)
}
