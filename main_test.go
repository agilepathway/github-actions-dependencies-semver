package main_test

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/magefile/mage/mage"
)

const (
	MagefileVerbose = "MAGEFILE_VERBOSE"
)

// nolint: lll, funlen
func TestSemverLookup(t *testing.T) {
	tests := map[string]struct {
		dependency     string
		expectedStdout string
		expectedStderr string
	}{
		"First example - make me more descriptive in the future": {"hadolint/hadolint", "v1.22.1", ""},
	}
	for name, tc := range tests {
		tc := tc

		t.Run(name, func(t *testing.T) {
			if len(tc.expectedStderr) > 0 {
				tc.expectedStderr = "Error: " + tc.expectedStderr
			}

			exitCode, stderr, stdout := semver()

			if (len(tc.expectedStderr) > 0) && (exitCode == 0) {
				t.Fatalf("got exit code %v, err: %s", exitCode, stderr)
			}

			if (len(tc.expectedStderr) == 0) && (exitCode != 0) {
				t.Fatalf("got exit code %v, err: %s", exitCode, stderr)
			}

			if actual := stdout.String(); actual != tc.expectedStdout {
				t.Fatalf("expected %q but got %q", tc.expectedStdout, actual)
			}

			if actual := stderr.String(); actual != tc.expectedStderr {
				t.Fatalf("expected %q but got %q", tc.expectedStderr, actual)
			}
		})
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Setenv(MagefileVerbose, "1") //nolint
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	//nolint
	defer func() {
		os.Unsetenv(MagefileVerbose)
	}()

	return m.Run()
}

func semver() (int, *bytes.Buffer, *bytes.Buffer) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	invocation := mage.Invocation{Stderr: stderr, Stdout: stdout}

	return mage.Invoke(invocation), stderr, stdout
}
