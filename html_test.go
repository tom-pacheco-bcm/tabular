package tabular_test

import (
	"bytes"
	"testing"

	"github.com/tom-pacheco-bcm/tabular"
)

func TestTableHtml(t *testing.T) {

	b := &bytes.Buffer{}

	table := []*struct {
		A string
		B int
		C float32
	}{
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
	}

	expect := `
<!DOCTYPE html>
<html lang="en-US">
  <head>
  </head>
  <body>
    <table>
      <thead>
        <tr><th> A </th><th class="number"> B </th><th class="number"> C </th></tr>
      </thead>
      <tbody>
        <tr><td> One </td><td class="number"> 1 </td><td class="number"> 1.000 </td></tr>
        <tr><td> Two </td><td class="number"> 2 </td><td class="number"> 2.000 </td></tr>
        <tr><td> Three </td><td class="number"> 3 </td><td class="number"> 3.000 </td></tr>
      </tbody>
      <tfoot>
      </tfoot>
    </table>
  </body>
</html>
`

	tb := tabular.From(table)
	w := tb.HtmlWriter()
	_, err := w.WriteTo(b)
	if err != nil {
		t.Error("failed to write html file.", err)
	}
	if b.String() != expect {
		t.Log(b.String())
		t.Error("expected html to match")
	}
}
