package tabular_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
)

func TestTable(t *testing.T) {

	b := &bytes.Buffer{}

	type ABC struct {
		A string
		B int
		C float32
	}

	table := []ABC{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	expect := []string{
		"A      B         C",
		"-----  -  --------",
		"One    1  1.000000",
		"Two    2  2.000000",
		"Three  3  3.000000",
		"",
	}

	tb := tabular.From(table)

	tw := tb.TextWriter()
	_, err := tw.WriteTo(b)
	if err != nil {
		t.Error("failed to write text table.")
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

func TestTableRef(t *testing.T) {

	b := &bytes.Buffer{}

	type ABC struct {
		A string
		B int
		C float32
	}

	table := []*ABC{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	expect := []string{
		"A      B         C",
		"-----  -  --------",
		"One    1  1.000000",
		"Two    2  2.000000",
		"Three  3  3.000000",
		"",
	}

	tb := tabular.From(table)
	tw := tb.TextWriter()
	_, err := tw.WriteTo(b)
	if err != nil {
		t.Error("failed to write text table.")
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

func TestTableRefRef(t *testing.T) {

	b := &bytes.Buffer{}

	type ABC struct {
		A string
		B *int
		C float32
	}
	var x int = 1

	table := []*ABC{
		{"One", &x, 1},
		{"Two", &x, 2},
		{"Three", &x, 3},
	}

	expect := []string{
		"A      B         C",
		"-----  -  --------",
		"One    1  1.000000",
		"Two    1  2.000000",
		"Three  1  3.000000",
		"",
	}

	tb := tabular.From(table)
	tw := tb.TextWriter()
	_, err := tw.WriteTo(b)
	if err != nil {
		t.Error("failed to write text table.")
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

type BorderFuncHorizontal func(int, string) string

type StringFunc func() string

type SsssI interface {
	string | BorderFuncHorizontal | func() string
}

type Sd interface {
	View() string
}

type Position float32

type BorderDecorator struct {
	dec interface{}
	p   Position
}

func NewBorderDecorator[T interface {
	string | BorderFuncHorizontal | func() string
}](p Position, dec T) BorderDecorator {

	return BorderDecorator{
		p:   p,
		dec: dec,
	}

}
