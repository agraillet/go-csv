// 参考Python的DictReader
package goCsv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type MapReader struct {
	Reader     *csv.Reader
	fieldnames []string
}

func NewMapReader(r io.Reader) *MapReader {
	return &MapReader{
		Reader: csv.NewReader(r),
	}
}

// Init 配置csv的基本参数
func (r *MapReader) Init(params csv.Reader) {
	r.Reader.Comma = params.Comma
	r.Reader.LazyQuotes = params.LazyQuotes
	r.Reader.TrimLeadingSpace = params.TrimLeadingSpace
}

// SetFieldnames 指定csv文件的字段名
// 如果不指定的话，则默认使用csv文件的第一行作为字段名
func (r *MapReader) SetFieldnames(fieldnames []string) {
	r.fieldnames = fieldnames
}

func (r *MapReader) GetFieldnames() (fieldnames []string, err error) {
	if len(r.fieldnames) == 0 {
		// 如果没有设置字段名，则默认为csv的第一行为字段名
		if r.fieldnames, err = r.Reader.Read(); err != nil {
			return nil, err
		}
	}
	return r.fieldnames, nil
}

// Read 读取一行记录
func (r *MapReader) Read() (record map[string]string, err error) {
	if len(r.fieldnames) == 0 {
		// 如果没有设置字段名，则默认为csv的第一行为字段名
		r.fieldnames, err = r.Reader.Read()
		if err != nil {
			return nil, err
		}
	}

	var rawRecord []string
	rawRecord, err = r.Reader.Read()
	if err != nil {
		return nil, err
	}

	length := len(r.fieldnames)
	record = make(map[string]string)
	for index := 0; index < length; index++ {
		field := r.fieldnames[index]
		if _, exists := record[field]; exists {
			return nil, fmt.Errorf("Multiple indices with the same name '%s'", field)
		}
		record[field] = rawRecord[index]
	}
	return record, err
}

// ReadAll 读取全部的内容
func (r *MapReader) ReadAll() (records []map[string]string, err error) {
	var record map[string]string
	for record, err = r.Read(); err == nil; record, err = r.Read() {
		records = append(records, record)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return records, nil
}
