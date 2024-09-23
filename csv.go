package tabular

import (
	"bytes"
	"encoding/csv"
	"io"
)

// CSV is a csv file writer
type CSV[T any] struct {
	Table[T]
}

// CSVWriter returns a writer that writes CSV formatted data
func (tbl *Table[T]) CSVWriter() *CSV[T] {
	return &CSV[T]{*tbl}
}

func (ct *CSV[T]) WriteTo(w io.Writer) (int64, error) {
	table := ct.Rows()

	b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(b)
	err := csvWriter.Write(ct.HeaderNames())
	if err != nil {
		return 0, err
	}
	err = csvWriter.WriteAll(table)
	if err != nil {
		return 0, err
	}
	return io.Copy(w, b)
}
