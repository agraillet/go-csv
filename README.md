# go-csv
实现类似Python的csv.DictReader和csv.DictWriter

## Example

```go
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

	writer.Flush()  // 注意最后需要刷到磁盘
}
```

## MapReader API

```go
func NewMapReader(r io.Reader) *MapReader

// Init 配置csv的基本参数
func (r *MapReader) Init(params csv.Reader)

// SetFieldnames 指定csv文件的字段名
// 如果不指定的话，则默认使用csv文件的第一行作为字段名
func (r *MapReader) SetFieldnames(fieldnames []string)

func (r *MapReader) GetFieldnames() (fieldnames []string, err error) 

// Read 读取一行记录
func (r *MapReader) Read() (record map[string]string, err error) 

// ReadAll 读取全部的内容
func (r *MapReader) ReadAll() (records []map[string]string, err error)
```

## MapWriter API

```go
func NewMapWriter(w io.Writer, fieldnames []string) *MapWriter 

// Init 配置csv的基本参数
func (w *MapWriter) Init(params csv.Writer) 

func (w *MapWriter) WriteHeader() (err error) 

func (w *MapWriter) WriteRow(row map[string]string) (err error) 

func (w *MapWriter) WriteRows(rows []map[string]string) (err error) 

// Flush 将数据刷到磁盘
func (w *MapWriter) Flush() 
```
