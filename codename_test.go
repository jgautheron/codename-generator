package codename_test

import (
	"testing"

	"github.com/jgautheron/codename-generator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCodenameGeneration(t *testing.T) {
	var cn string
	var err error

	Convey("Codename should be generated", t, func() {
		cn, err = codename.Get(codename.Sanitized)
		So(err, ShouldBeNil)
		So(len(cn), ShouldBeGreaterThan, 3)

		cn, err = codename.Get(codename.SpacedString)
		So(err, ShouldBeNil)
		So(len(cn), ShouldBeGreaterThan, 3)
	})
}
