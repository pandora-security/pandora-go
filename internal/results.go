/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

type PlainText string
type CipherText string

type DecryptionResult struct {
	Decrypted	PlainText	`json:"decrypted"`
}

type EncryptionResult struct {
	Encrypted	CipherText	`json:"encrypted"`
}
