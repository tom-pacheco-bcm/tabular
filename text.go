package tabular

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// Text is a plain text writer
type Text[T any] struct {
	Table[T]
}

// TextWriter returns a writer that writes plain text tables
func (tbl *Table[T]) TextWriter() *Text[T] {
	return &Text[T]{Table: *tbl}
}

func (tt *Text[T]) WriteTo(dst io.Writer) (int64, error) {

	var err error

	table := tt.Rows()

	formats := make([]string, len(tt.Columns))

	// update the formatting for column widths

	tt.autoWidth(table)

	for i := range tt.Columns {
		formats[i] = fmt.Sprintf("%%-%ds", tt.Columns[i].Width)
	}

	// right justify all numeric types
	// everything else is left justified
	for i := range tt.Columns {
		switch tt.Columns[i].FieldType {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64, reflect.Int32,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Uint32, reflect.Float32, reflect.Float64:
			tt.Columns[i].HeaderFormat = fmt.Sprintf("%%%ds", tt.Columns[i].Width)
		default:
			tt.Columns[i].HeaderFormat = fmt.Sprintf("%%-%ds", tt.Columns[i].Width)
		}
	}

	b := &bytes.Buffer{}

	// header row

	for i, col := range tt.Columns {
		if i > 0 {
			_, err = fmt.Fprint(b, "  ")
			if err != nil {
				return 0, err
			}
		}
		_, err = fmt.Fprint(b, col.Header())
		if err != nil {
			return 0, err
		}
	}

	_, err = fmt.Fprintln(b)
	if err != nil {
		return 0, err
	}

	// header separator row

	for i, col := range tt.Columns {
		if i > 0 {
			_, err = fmt.Fprint(b, "  ")
			if err != nil {
				return 0, err
			}
		}
		_, err = fmt.Fprint(b, strings.Repeat("-", col.Width))
		if err != nil {
			return 0, err
		}
	}

	_, err = fmt.Fprintln(b)
	if err != nil {
		return 0, err
	}

	// the data table

	for _, row := range table {
		for i := range tt.Columns {
			if i > 0 {
				_, err = fmt.Fprint(b, "  ")
				if err != nil {
					return 0, err
				}
			}
			_, err = fmt.Fprintf(b, formats[i], row[i])
			if err != nil {
				return 0, err
			}
		}
		_, err = fmt.Fprintln(b)
		if err != nil {
			return 0, err
		}
	}

	return b.WriteTo(dst)
}

func (tt *Text[T]) autoWidth(table [][]string) {

	for i := range table {
		row := table[i]
		for j := range tt.Columns {
			if tt.Columns[j].Width < len(row[j]) {
				tt.Columns[j].Width = len(row[j])
			}
		}
	}

	for i := range tt.Columns {
		if tt.Columns[i].Width < len(tt.Columns[i].HeaderName) {
			tt.Columns[i].Width = len(tt.Columns[i].HeaderName)
		}
	}

}
