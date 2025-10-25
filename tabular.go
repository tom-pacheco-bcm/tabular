// tabular is a simple table generation and formatting utility.
//
// Tabular reads a slice and can generate a table based on the public properties
// of the data type. It can then generate text, CSV, or xlsx outputs.
//
// # Example
//
//	func main() {
//	    dirList, err := os.ReadDir("/")
//	    if err != nil {
//	        os.Exit(1)
//	    }
//
//	    dirs := make([]struct {
//	        Name    string
//	        IsDir bool
//	    }, len(dirList))
//
//	    for i, d := range dirList {
//	        dirs[i].Name = d.Name()
//	        dirs[i].IsDir = d.IsDir()
//	    }
//
//	    dt := tabular.From(dirs)
//	    w := dt.TextWriter()
//	    w.WriteTo(os.Stdout)
//	}
//
// output:
//
//	Name                          IsDir
//	----------------------------  -----
//	$RECYCLE.BIN                   true
//	Data                           true
//	README                        false
package tabular

import (
	"fmt"
	"reflect"
)

// Column represents a table column
type Column struct {
	FieldIndex   int
	FieldType    reflect.Kind
	FieldName    string
	HeaderName   string
	HeaderFormat string
	Format       string
	Width        int
	Hidden       bool
}

func (c *Column) Header() string {
	return fmt.Sprintf(c.HeaderFormat, c.HeaderName)
}

// Table
type Table[T any] struct {
	Columns []Column
	data    []T
	Footer  bool
}

// From creates a table from a array of some data type
func From[T any](data []T) *Table[T] {
	return &Table[T]{
		Columns: columns(data),
		data:    data,
	}
}

func columns(s any) []Column {
	val := reflect.ValueOf(s)
	if !val.IsValid() {
		return nil
	}

	for val.Kind() == reflect.Interface || val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	typ := val.Type()

	if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		cols := make([]Column, 0, n)
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if !f.IsExported() {
				continue
			}
			cols = append(cols,
				Column{
					FieldName:    f.Name,
					FieldIndex:   i,
					FieldType:    f.Type.Kind(),
					HeaderName:   f.Name,
					HeaderFormat: "%s",
					Width:        -1,
					Format:       getFormat(f.Type),
				})
		}
		return cols
	}
	return nil
}

func getFormat(t reflect.Type) string {
	k := t.Kind()
	if k == reflect.Pointer {
		k = t.Elem().Kind()
	}
	switch k {
	case reflect.Bool:
		return "%t"
	case reflect.String:
		return "%s"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64, reflect.Int32:
		return "%d"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Uint32:
		return "%d"
	case reflect.Float32, reflect.Float64:
		return "%f"
	default:
		return "%v"
	}
}

func (tbl *Table[T]) HeaderNames() []string {
	cols := tbl.visibleColumns()
	row := make([]string, len(cols))
	for i := range cols {
		row[i] = cols[i].HeaderName
	}
	return row
}

func (tbl *Table[T]) visibleColumns() []Column {
	cols := make([]Column, 0, len(tbl.Columns))
	for i := range tbl.Columns {
		if tbl.Columns[i].Hidden {
			continue
		}
		cols = append(cols, tbl.Columns[i])
	}
	return cols
}

// Rows generates a string table from the typed array data
func (tbl *Table[T]) Rows() [][]string {
	values := reflect.ValueOf(tbl.data)
	cols := tbl.visibleColumns()
	colCount := len(cols)
	table := make([][]string, 0, len(tbl.data))
	for i := range tbl.data {
		row := make([]string, colCount)
		table = append(table, row)
		v := values.Index(i)
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		for j, col := range cols {
			f := v.Field(col.FieldIndex)
			if f.Kind() == reflect.Pointer {
				f = f.Elem()
			}
			row[j] = fmt.Sprintf(col.Format, f.Interface())
		}
	}
	return table
}
