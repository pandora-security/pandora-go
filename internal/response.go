/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

type Response struct {
	Success		bool			`json:"success"`
	Initiator	*Application	`json:"initiator,omitempty"`
	Message		string			`json:"message,omitempty"`
	Data		string			`json:"data,omitempty"`
	Signature	string			`json:"signature,omitempty"`
}