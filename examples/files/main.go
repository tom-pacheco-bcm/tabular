package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/tom-pacheco-bcm/tabular"
)

type dirInfo struct {
	Name     string
	IsDir    bool
	Size     int64
	Modified string
}

func readDirEntry(d fs.DirEntry) dirInfo {
	di := dirInfo{}
	di.Name = d.Name()
	di.IsDir = d.IsDir()
	info, err := d.Info()
	if err != nil {
		return di
	}
	di.Size = info.Size()
	di.Modified = info.ModTime().Format("01/02/2006 15:04:05")
	return di
}

func main() {

	dirList, err := os.ReadDir("/")
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}

	dirs := make([]dirInfo, 0, len(dirList))

	for _, d := range dirList {
		di := readDirEntry(d)
		dirs = append(dirs, di)
	}

	dt := tabular.From(dirs[:10])

	tw := dt.TextWriter()
	tw.WriteTo(os.Stdout)

	fmt.Println()

	cw := dt.CSVWriter()
	_, err = cw.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println("error", err)
	}

}
