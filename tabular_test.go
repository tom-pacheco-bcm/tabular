package tabular_test

import (
	"strconv"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
)

func TestTable1(t *testing.T) {

	type ABC struct {
		A string
		B int
		C float32
	}

	headers := []string{
		"A",
		"B",
		"C",
	}

	formats := []string{
		"%s",
		"%d",
		"%f",
	}

	table := []ABC{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	tb := tabular.From(table)

	for i, name := range tb.HeaderNames() {
		if name != headers[i] {
			t.Errorf("expected header %d %s got %s", i, headers[i], name)
		}
	}

	for i, r := range tb.Rows() {
		er := table[i]
		if r[0] != er.A {
			t.Errorf("expected row[%d].A %s got %s", i, er.A, r[0])
		}
		if s := strconv.Itoa(er.B); r[1] != s {
			t.Errorf("expected row[%d].B %s got %s", i, s, r[1])
		}
		if s := strconv.FormatFloat(float64(er.C), 'f', 6, 32); r[2] != s {
			t.Errorf("expected row[%d].C %s got %s", i, s, r[2])
		}
	}

	for i, c := range tb.Columns {
		if c.FieldIndex != i {
			t.Errorf("expected Columns[%d].FieldIndex %d got %d", i, c.FieldIndex, i)
		}
		if c.FieldName != c.HeaderName {
			t.Errorf("expected %d FieldName and HeaderName to be the same %s vs %s", i, c.FieldName, c.HeaderName)
		}
		if c.FieldName != headers[i] {
			t.Errorf("expected %d FieldName and HeaderName to be the same %s vs %s", i, c.FieldName, headers[i])
		}
		if c.Format != formats[i] {
			t.Errorf("expected %d Format %s got %s", i, c.Format, formats[i])
		}
	}

}
