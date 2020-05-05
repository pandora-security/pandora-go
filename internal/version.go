/**
 * This code is part of internal core of
 * Pandora Security Driver for Go
 **/

package internal

import "fmt"

type SemanticVersion struct {
	Major		int		`json:"major"`
	Minor		int		`json:"minor"`
	Patch		int		`json:"patch"`
}

func (version *SemanticVersion) String() string {
	return fmt.Sprintf("v%d.%d.%d", version.Major, version.Minor, version.Patch)
}