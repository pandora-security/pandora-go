/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

import (
	"crypto/ed25519"
	"encoding/base64"
)

const publicKey = "Y/97DVIGeTgOzFW7j+vXS8g8j0UUcum3g50QfO1CH+c="

func Verify(data string, signature string) (isAuthentic bool, err error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, Base64DecodingError
	}
	dataBytes := []byte(data)
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, Base64DecodingError
	}
	return ed25519.Verify(publicKeyBytes, dataBytes, signatureBytes), nil
}
