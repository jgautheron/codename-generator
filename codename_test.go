package codename_test

import (
	"testing"

	"github.com/jgautheron/codename-generator"
)

func TestCodenameGeneration(t *testing.T) {
	var cn string
	var err error

	cn, err = codename.Get(codename.Sanitized)
	if err != nil {
		t.Error("The codename was not generated")
	}
	if len(cn) < 3 {
		t.Error("The codename generated is invalid")
	}

	cn, err = codename.Get(codename.SpacedString)
	if err != nil {
		t.Error("The codename was not generated")
	}
	if len(cn) < 3 {
		t.Error("The codename generated is invalid")
	}
}
