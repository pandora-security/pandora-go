/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

type Command struct {
	Command			string		`json:"command"`
	RequiresAuth	bool		`json:"requiresAuth"`
}

var DecryptCommand = Command{
	Command:      "decrypt",
	RequiresAuth: true,
}

var EncryptCommand = Command{
	Command:      "encrypt",
	RequiresAuth: true,
}

var VersionCommand = Command{
	Command:      "version",
	RequiresAuth: false,
}
