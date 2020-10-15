package coreutils

import (
	"testing"
)

func TestLoadAvg(t *testing.T) {
	_, err := LoadAvg()
	if err != nil {
		t.Fatal(err)
	}
}
