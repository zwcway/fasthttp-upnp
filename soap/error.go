package soap

import (
	"encoding/xml"
	"fmt"
)

type Error struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:control-1-0 Error"`
	Code    uint     `xml:"errorCode"`
	Desc    string   `xml:"errorDescription"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Desc)
}

const (
	NoErrorCode                   = 200
	BadRequestErrorCode           = 400
	InvalidActionErrorCode        = 401
	ActionFailedErrorCode         = 501
	ArgumentValueInvalidErrorCode = 600
)

var (
	InvalidActionError        = NewErrorf(401, "Invalid Action")
	ArgumentValueInvalidError = NewErrorf(600, "The argument value is invalid")
)

func NewErrorf(code uint, tpl string, args ...interface{}) *Error {
	return &Error{Code: code, Desc: fmt.Sprintf(tpl, args...)}
}

func NewFailed(err error) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	return NewErrorf(ActionFailedErrorCode, err.Error())
}

func NewError(code uint, err error) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	return NewErrorf(code, err.Error())
}
