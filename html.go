package tabular

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

// HTML is a html file writer
type HTML[T any] struct {
	Table[T]
	Css string
}

// HtmlWriter returns a writer that writes HTML formatted data
func (tbl *Table[T]) HtmlWriter() *HTML[T] {
	return &HTML[T]{
		Table: *tbl,
		Css:   "",
	}
}

// WriteTo writes the data to the given writer (implements io.WriterTo).
func (ct *HTML[T]) WriteTo(w io.Writer) (int64, error) {
	return ct.WriteHtml(w)
}

// WriteHtml writes the data as a html document
func (ct *HTML[T]) WriteHtml(dst io.Writer) (int64, error) {

	var err error

	table := ct.Rows()

	formats := make([]string, len(ct.Columns))

	// right justify all numeric types
	// everything else is left justified
	for i := range ct.Columns {
		switch ct.Columns[i].FieldType {
		case reflect.Bool, reflect.Int, reflect.Uint,
			reflect.Int8, reflect.Int16, reflect.Int64, reflect.Int32,
			reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Uint32,
			reflect.Float32, reflect.Float64:
			ct.Columns[i].HeaderFormat = "<th class=\"number\"> %s </th>"
			formats[i] = "<td class=\"number\"> %s </td>"
		default:
			ct.Columns[i].HeaderFormat = "<th> %s </th>"
			formats[i] = "<td> %s </td>"
		}
	}

	b := &bytes.Buffer{}

	fmt.Fprint(b, htmlHeader)

	// header row

	fmt.Fprintln(b, "    <table>")
	fmt.Fprintln(b, "      <thead>")
	fmt.Fprint(b, "        <tr>")

	for _, col := range ct.Columns {
		_, err = fmt.Fprint(b, col.Header())
		if err != nil {
			return 0, err
		}
	}

	fmt.Fprintln(b, "</tr>")
	fmt.Fprintln(b, "      </thead>")

	// the data table
	n := len(table)
	if ct.Footer {
		n--
	}

	fmt.Fprintln(b, "      <tbody>")

	for _, row := range table[:n] {
		fmt.Fprint(b, "        <tr>")
		for i := range ct.Columns {
			_, err = fmt.Fprintf(b, formats[i], row[i])
			if err != nil {
				return 0, err
			}
		}
		fmt.Fprintln(b, "</tr>")
	}

	fmt.Fprintln(b, "      </tbody>")
	fmt.Fprintln(b, "      <tfoot>")
	if ct.Footer {
		row := table[n]
		for i := range ct.Columns {
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

	fmt.Fprintln(b, "      </tfoot>")
	fmt.Fprintln(b, "    </table>")
	fmt.Fprint(b, htmlFooter)
	return b.WriteTo(dst)
}

const htmlHeader = `
<!DOCTYPE html>
<html lang="en-US">
  <head>
  </head>
  <body>
`
const htmlFooter = `  </body>
</html>
`
