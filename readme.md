# tabular

Tabular is a module for simple table outputs with options for csv or plain text.

## Usage

Download the package:

	go get github.com/tom-pacheco-bcm/tabular

### Example

```go
package main

import (
    "fmt"
    "io/fs"
    "os"

    "github.com/tom-pacheco-bcm/tabular"
)

func main() {
    dirList, err := os.ReadDir("/")
    if err != nil {
        os.Exit(1)
    }

    dirs := make([]struct {
        Name  string
        IsDir bool
    }, len(dirList))

    for i, d := range dirList {
        dirs[i].Name = d.Name()
        dirs[i].IsDir = d.IsDir()
    }

    dt := tabular.From(dirs)
    w := dt.TextWriter()
    w.WriteTo(os.Stdout)
}
```

_output:_
```
Name                          IsDir
----------------------------  -----
$RECYCLE.BIN                   true
Data                           true
README                        false
```
