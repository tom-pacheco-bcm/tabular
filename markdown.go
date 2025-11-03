package tabular

import (
	"bytes"
	"fmt"
	"io"
)

// Markdown is a markdown file writer
type Markdown[T any] struct {
	Table[T]
}

// HtmlWriter returns a writer that writes Markdown formatted data
func (tbl *Table[T]) MarkdownWriter() *Markdown[T] {
	return &Markdown[T]{*tbl}
}

// WriteTo writes the data to the given writer (implements io.WriterTo).
func (md *Markdown[T]) WriteTo(w io.Writer) (int64, error) {
	return md.WriteMarkdown(w)
}

// WriteMarkdown writes the data as a html document
func (md *Markdown[T]) WriteMarkdown(dst io.Writer) (int64, error) {

	var err error

	table := md.Rows()

	formats := make([]string, len(md.Columns))

	for i := range md.Columns {
		formats[i] = "%s"
		md.Columns[i].HeaderFormat = "%s"
	}

	b := &bytes.Buffer{}
	fmt.Fprintln(b)

	// header row

	for _, col := range md.Columns {
		fmt.Fprint(b, " | ")
		_, err = fmt.Fprint(b, col.Header())
		if err != nil {
			return 0, err
		}
	}
	fmt.Fprintln(b, " |")

	for range md.Columns {
		fmt.Fprint(b, " | -")
	}
	fmt.Fprintln(b, " |")

	// the data table

	n := len(table)
	if md.Footer {
		n--
	}

	for _, row := range table[:n] {
		for i := range md.Columns {
			fmt.Fprint(b, " | ")
			_, err = fmt.Fprintf(b, formats[i], row[i])
			if err != nil {
				return 0, err
			}
		}
		fmt.Fprintln(b, " |")
	}
	if md.Footer {
		row := table[n]
		for i := range md.Columns {
			fmt.Fprint(b, " |")
			_, err = fmt.Fprintf(b, formats[i], row[i])
			if err != nil {
				return 0, err
			}
			fmt.Fprintln(b, " |")
		}
	}

	return b.WriteTo(dst)
}
