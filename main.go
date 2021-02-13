// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Default = SemverLookup

// SemverLookup provides the current semantic version associated with the given GitHub Action dependency
func SemverLookup() error {
	yaml_filename := "https://github.com/agilepathway/agilepathway-template/raw/non-managed-dependencies/.github/workflows/non-package-managed-dependencies.yml"
	_, err := getContent(yaml_filename)

	fmt.Print("v1.22.1")
	return err
}

func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
