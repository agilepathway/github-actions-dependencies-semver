// +build mage

package main

import "fmt"

var Default = SemverLookup

// SemverLookup provides the current semantic version associated with the given GitHub Action dependency
func SemverLookup() error {
	// https://github.com/agilepathway/agilepathway-template/raw/non-managed-dependencies/.github/workflows/non-package-managed-dependencies.yml
	fmt.Print("v1.22.1")
	return nil
}
