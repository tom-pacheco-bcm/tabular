package tabular_test

import (
	"os"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
	"github.com/xuri/excelize/v2"
)

func TestTableXlsx(t *testing.T) {
	table := []*struct {
		A string
		B int
		C float32
	}{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	expect := [][]string{
		{"A", "B", "C"},
		{"One", "1", "1.000000"},
		{"Two", "2", "2.000000"},
		{"Three", "3", "3.000000"},
	}

	testFile := "./test.xlsx"

	tb := tabular.From(table)
	w := tb.XLSXWriter()
	err := w.WriteToFile(testFile)
	if err != nil {
		t.Error("failed to write xlsx file.", err)
	}

	f, err := excelize.OpenFile(testFile)
	if err != nil {
		t.Error("failed to open xlsx file.", err)
	}
	rows, err := f.GetRows(f.GetSheetName(f.GetActiveSheetIndex()))
	if err != nil {
		t.Error("failed to read xlsx file.", err)
	}
	for i, row := range rows {
		if i >= len(expect) {
			break
		}
		for j, cell := range row {
			if i >= len(expect[i]) {
				break
			}
			if cell != expect[i][j] {
				t.Errorf("cell %d,%d expected %q got %q", i, j, expect[i][j], cell)
			}
		}
	}
	os.Remove(testFile)
}
