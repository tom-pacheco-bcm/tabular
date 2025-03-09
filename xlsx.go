package tabular

import (
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

// XLSX is a xlsx file writer
type XLSX[T any] struct {
	Table[T]
}

// CSVWriter returns a writer that writes CSV formatted data
func (tbl *Table[T]) XLSXWriter() *XLSX[T] {
	return &XLSX[T]{*tbl}
}

func (ct *XLSX[T]) WriteToFile(name string) error {

	f := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	n := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(n)
	{
		header := ct.HeaderNames()
		f.SetSheetRow(sheetName, "A1", &header)
	}

	for i, row := range ct.Rows() {
		f.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+2), &row)
	}

	err := f.SaveAs(name)
	if err != nil {
		return err
	}
	return nil
}

func (ct *XLSX[T]) WriteTo(w io.Writer) (int64, error) {
	f := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	n := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(n)
	{
		header := ct.HeaderNames()
		f.SetSheetRow(sheetName, "A1", &header)
	}

	for i, row := range ct.Rows() {
		f.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+2), &row)
	}
	return f.WriteTo(w)
}
