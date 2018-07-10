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

// SetSkip 设置跳过前面的若干条记录
func (r *MapReader) SetSkip(skip int) 

// SetLimit 设置只提取若干条记录
func (r *MapReader) SetLimit(limit int) 

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

// NewMapWriterSimple 简化的writer对象，可以通过SetHeader方法来设置Header
func NewMapWriterSimple(w io.Writer) *MapWriter 

func (w *MapWriter) SetHeader(fieldnames []string) 

// SetFieldNotSetErr 字段未设置时，是否报错
// 默认为false，即当字段为设置时，会自动使用空字符串补充
func (w *MapWriter) SetFieldNotSetErr(fieldNotSetErr bool) 

// Init 配置csv的基本参数
func (w *MapWriter) Init(params csv.Writer) 

func (w *MapWriter) WriteHeader() (err error) 

func (w *MapWriter) WriteRow(row map[string]string) (err error) 

func (w *MapWriter) WriteRows(rows []map[string]string) (err error) 

// Flush 将数据刷到磁盘
func (w *MapWriter) Flush() 
```

## Reader API

```go
func NewReader(r io.Reader) *Reader

// Init 配置csv的基本参数
func (r *Reader) Init(params csv.Reader)

// SetSkip 设置跳过前面的若干条记录
func (r *Reader) SetSkip(skip int) 

// SetLimit 设置只提取若干条记录
func (r *Reader) SetLimit(limit int) 

func (r *Reader) GetFieldnames() (fieldnames []string, err error) 

// Read 读取一行记录
func (r *Reader) Read() (record []string, err error) 

// ReadAll 读取全部的内容
func (r *Reader) ReadAll() (records [][]string, err error)
```

## utils

```go
// CountLines 统计csv文件的记录数
func CountLines(r io.Reader) (n int, err error)

```
