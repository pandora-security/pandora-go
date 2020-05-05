/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

import (
	"encoding/base64"
	"encoding/json"
)

func DecodeResponse(src string) (*Response, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, Base64DecodingError
	}
	var dst Response
	if err = json.Unmarshal(data, &dst); err != nil {
		return nil, JsonMarshalingError
	}
	return &dst, nil
}

func DecodeDecryptionResult(src string) (*DecryptionResult, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, Base64DecodingError
	}
	var dst DecryptionResult
	if err = json.Unmarshal(data, &dst); err != nil {
		return nil, JsonMarshalingError
	}
	return &dst, nil
}

func DecodeEncryptionResult(src string) (*EncryptionResult, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, Base64DecodingError
	}
	var dst EncryptionResult
	if err = json.Unmarshal(data, &dst); err != nil {
		return nil, JsonMarshalingError
	}
	return &dst, nil
}

func DecodeVersionResult(src string) (*SemanticVersion, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, Base64DecodingError
	}
	var dst SemanticVersion
	if err = json.Unmarshal(data, &dst); err != nil {
		return nil, JsonMarshalingError
	}
	return &dst, nil
}

func EncodePlainText(src PlainText) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func DecodePlainText(src PlainText) (PlainText, error) {
	dst, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return "", Base64DecodingError
	}
	return PlainText(dst), nil
}