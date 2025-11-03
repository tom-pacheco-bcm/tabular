package tabular_test

import (
	"bytes"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
)

func TestTableMarkdown(t *testing.T) {

	b := &bytes.Buffer{}

	table := []*struct {
		A string
		B int
		C float32
	}{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	expect := `
 | A | B | C |
 | - | - | - |
 | One | 1 | 1.000 |
 | Two | 2 | 2.000 |
 | Three | 3 | 3.000 |
`

	tb := tabular.From(table)
	w := tb.MarkdownWriter()
	_, err := w.WriteTo(b)
	if err != nil {
		t.Error("failed to write html file.", err)
	}
	if b.String() != expect {
		t.Log(b.String())
		t.Error("expected html to match")
	}
}
