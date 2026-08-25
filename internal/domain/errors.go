package domain

import "errors"

var ErrInvalid = errors.New("invalid accounting record")
var ErrNotFound = errors.New("record not found")
var ErrConflict = errors.New("record conflict")

func IsMissing(err error) bool { return errors.Is(err, ErrNotFound) }
func IsInvalid(err error) bool { return errors.Is(err, ErrInvalid) }
