package box_test

import (
	"testing"

	"github.com/logrusorgru/aurora"

	"github.com/altipla-consulting/box"
)

func ExampleBox() {
	var o box.Box
	o.AddLine("foo", "bar")
	o.AddLine("before", aurora.Red("colored"), "after")
	o.Render()
}

func TestBox(t *testing.T) {
	var o box.Box
	o.AddLine("foo", "bar")
	o.AddLine("before", aurora.Red("colored"), "after")
	o.Render()
}
