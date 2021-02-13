// +build mage

package main

import "fmt"

var Default = SemverLookup

// SemverLookup provides the current semantic version associated with the given GitHub Action dependency
func SemverLookup() error {
	fmt.Print("v1.22.1")
	return nil
}
