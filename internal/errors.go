/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

import "fmt"

type ErrorCode int

type Error struct {
	Code		ErrorCode	`json:"code"`
	Message		string		`json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s [code: %d]", e.Message, e.Code)
}

var PandoraNotInstalledError = Error{
	Code:    100,
	Message: "pandora is not installed",
}

var PandoraRuntimeError = Error{
	Code:    101,
	Message: "pandora runtime error",
}

var Base64DecodingError = Error{
	Code:    200,
	Message: "base64 decoding error",
}

var JsonMarshalingError = Error{
	Code:    201,
	Message: "json marshalling error",
}

var InvalidSignatureError = Error{
	Code:    300,
	Message: "invalid signature",
}