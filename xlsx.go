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

// WriteToFile writes the data to the given file
// If the file already exists, it will be overwritten
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

	err := ct.WriteToSheet(f, sheetName)
	if err != nil {
		return err

	}
	err = f.SaveAs(name)
	if err != nil {
		return err
	}
	return nil
}

// WriteTo writes the data to the given writer (implements io.WriterTo).
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

	ct.WriteToSheet(f, sheetName)

	return f.WriteTo(w)
}

// WriteToSheet writes the data to the given sheet
// If the sheet already exists, it may be overwritten
func (ct *XLSX[T]) WriteToSheet(f *excelize.File, name string) error {

	header := ct.HeaderNames()
	err := f.SetSheetRow(name, "A1", &header)
	if err != nil {
		return err
	}
	for i, row := range ct.Rows() {
		err = f.SetSheetRow(name, fmt.Sprintf("A%d", i+2), &row)
		if err != nil {
			return err
		}
	}
	return nil
}
