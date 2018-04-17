package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ibbd-dev/go-csv"
)

func main() {
	fname := "./test.csv"
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := goCsv.NewMapReader(f)
	fieldnames, err := reader.GetFieldnames()
	if err != nil {
		panic(err)
	}

	outFilename := "./out.csv"
	wf, err := os.Create(outFilename)
	if err != nil {
		panic(err)
	}
	defer wf.Close()
	writer := goCsv.NewMapWriter(wf, fieldnames)
	writer.WriteHeader()

	for {
		row, err := reader.Read()
		if err == io.EOF {
			fmt.Printf("the file is over")
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Row 1: %+v\n", row)

		if err = writer.WriteRow(row); err != nil {
			panic(err)
		}
	}

	writer.Flush() // 注意最后需要刷到磁盘
}
