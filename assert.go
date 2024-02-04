package testutil

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertEqual(t *testing.T, description string, actual interface{}, expected interface{}, cmpOpts ...cmp.Option) {
	if !cmp.Equal(actual, expected, cmpOpts...) {
		t.Errorf("Unexpected %s:\n%s\n\n", description, cmp.Diff(expected, actual, cmpOpts...))
	}
}

func AssertErrorContains(t *testing.T, err error, pieces ...string) {
	for _, p := range pieces {
		if err == nil {
			if p != "" {
				t.Errorf("Missing error string `%s`\nError: <nil>", p)
			}
		} else if !strings.Contains(err.Error(), p) {
			t.Errorf("Missing error string `%s`\nError: %s", p, err.Error())
		}
	}
}
