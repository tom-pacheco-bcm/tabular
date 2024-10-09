package tabular_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
)

func TestTableCsv(t *testing.T) {

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

	expect := []string{
		"A,B,C",
		"One,1,1.000000",
		"Two,2,2.000000",
		"Three,3,3.000000",
		"",
	}

	tb := tabular.From(table)
	w := tb.CSVWriter()
	_, err := w.WriteTo(b)
	if err != nil {
		t.Error("failed to write csv file.")
	}
	r := b.String()
	lines := strings.Split(r, "\n")
	if len(lines) != len(expect) {
		t.Errorf("expected %d lines got %d", len(expect), len(lines))
	}
	for i := range lines {
		if lines[i] != expect[i] {
			t.Errorf("line %d expected %q got %q", i, expect[i], lines[i])
		}
	}
}
