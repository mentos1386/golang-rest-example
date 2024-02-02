// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Group) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.UserIds == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "user_ids",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
