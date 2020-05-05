/**
 * Pandora Security Driver for Go
 *
 * An official Go module which will help any Go application to communicate
 * with Pandora Security Box APIs.
 **/
package pandora

import (
	"errors"
	"fmt"
	"github.com/danang-id/pandora-go/internal"
)

// Pandora Security Driver
type Pandora struct {
	ApplicationGUID		string						`json:"applicationGuid"`
	MinimumVersion		*internal.SemanticVersion	`json:"minimumVersion"`
}

var pandoraExecutable = internal.NewExecutable("pandora")

func (pandora *Pandora) executeCommand(command internal.Command, args ...string) (*internal.Response, error) {
	if !pandoraExecutable.Found() {
		return nil, internal.PandoraNotInstalledError
	}
	arguments := []string{ command.Command }
	if command.RequiresAuth {
		arguments = append(arguments, "-a", pandora.ApplicationGUID)
	}
	arguments = append(arguments, args...)
	output, err := pandoraExecutable.Run(arguments...)
	if err != nil {
		return nil, err
	}
	response, err := internal.DecodeResponse(output)
	if err != nil {
		return nil, err
	}
	isSignatureValid, err := internal.Verify(response.Data, response.Signature)
	if err != nil {
		return response, err
	}
	if !isSignatureValid {
		return response, internal.InvalidSignatureError
	}
	return response, nil
}

// Check the version of Pandora installed on the system.
func (pandora *Pandora) CheckVersion() (*internal.SemanticVersion, error) {
	response, err := pandora.executeCommand(internal.VersionCommand)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(response.Message)
	}
	result, err := internal.DecodeVersionResult(response.Data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Decrypt a given cipher text into a plain text.
func (pandora *Pandora) Decrypt(cipher internal.CipherText) (*internal.Application, internal.PlainText, error) {
	isSatisfied,  err := pandora.IsSatisfyMinimumVersion()
	if !isSatisfied {
		return nil, "", err
	}
	response, err := pandora.executeCommand(internal.DecryptCommand, string(cipher))
	if err != nil {
		return nil, "", err
	}
	if !response.Success {
		return response.Initiator, "", errors.New(response.Message)
	}
	result, err := internal.DecodeDecryptionResult(response.Data)
	if err != nil {
		return response.Initiator, "", err
	}
	plain, err := internal.DecodePlainText(result.Decrypted)
	if err != nil {
		return response.Initiator, "", err
	}
	return response.Initiator, plain, nil
}

// Encrypt a given plain text into a ciphered text.
func (pandora *Pandora) Encrypt(plain internal.PlainText) (*internal.Application, internal.CipherText, error) {
	isSatisfied,  err := pandora.IsSatisfyMinimumVersion()
	if !isSatisfied {
		return nil, "", err
	}
	response, err := pandora.executeCommand(internal.EncryptCommand, internal.EncodePlainText(plain))
	if err != nil {
		return nil, "", err
	}
	if !response.Success {
		return response.Initiator, "", errors.New(response.Message)
	}
	result, err := internal.DecodeEncryptionResult(response.Data)
	if err != nil {
		return response.Initiator, "", err
	}
	return response.Initiator, result.Encrypted, nil
}

// Check whether Pandora is installed or not on the system.
func (pandora *Pandora) IsInstalled() bool {
	return pandoraExecutable.Found()
}

// Check whether the installed Pandora's version is satisfy the minimum
// version required.
func (pandora *Pandora) IsSatisfyMinimumVersion() (bool, error) {
	if pandora.MinimumVersion == nil {
		return true, nil
	}
	version, err := pandora.CheckVersion()
	if err != nil {
		return false, err
	} else {
		err = errors.New(fmt.Sprintf("required Pandora %s; installed %s", pandora.MinimumVersion, version))
	}
	if version.Major < pandora.MinimumVersion.Major {
		return false, err
	} else if version.Major == pandora.MinimumVersion.Major {
		if version.Minor < pandora.MinimumVersion.Minor {
			return false, err
		} else if version.Minor == pandora.MinimumVersion.Minor {
			if version.Patch < pandora.MinimumVersion.Patch {
				return false, err
			}
		}
	}
	return true, nil
}

// Set the minimum version of Pandora that required to be installed.
func (pandora *Pandora) SetMinimumVersion(version *internal.SemanticVersion) {
	pandora.MinimumVersion = version
}

// Create new object which represent Pandora for specific application.
func New(applicationGUID string) *Pandora {
	return &(Pandora{ApplicationGUID: applicationGUID})
}

// Checks if an error is indicating that Pandora was not installed
func IsPandoraNotInstalledError(err error) bool {
	return errors.Is(err, internal.PandoraNotInstalledError)
}


// Checks if an error is indicating that Pandora execution was failed
func IsPandoraRuntimeError(err error) bool {
	return errors.Is(err, internal.PandoraRuntimeError)
}

// Checks if an error is indicating that base64 decoding was failed
func IsBase64DecodingError(err error) bool {
	return errors.Is(err, internal.Base64DecodingError)
}

// Checks if an error is indicating that JSON marshalling was failed
func IsJsonMarshallingError(err error) bool {
	return errors.Is(err, internal.JsonMarshalingError)
}


// Checks if an error is indicating that the provided signature is
// invalid
func IsInvalidSignatureError(err error) bool {
	return errors.Is(err, internal.InvalidSignatureError)
}