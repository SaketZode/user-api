package errorparser

import (
	"errors"
	"user-api/resterrors"

	"gorm.io/gorm"
)

func ParseError(err error) *resterrors.RestError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resterrors.NewNotFoundError("Record with given parameters not found!")
	}
	return resterrors.NewInternalServerError("Something went wrong!")
}
