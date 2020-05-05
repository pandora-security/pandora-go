/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

type Application struct {
	Guid 	string	`json:"guid,omitempty"`
	Name 	string	`json:"name,omitempty"`
	Author 	string	`json:"author,omitempty"`
}
