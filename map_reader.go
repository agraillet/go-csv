// 参考Python的DictReader
package goCsv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type MapReader struct {
	Reader
}

func NewMapReader(r io.Reader) *MapReader {
	reader := &MapReader{}
	reader.Reader.Reader = csv.NewReader(r)
	return reader
}

// Read 读取一行记录
func (r *MapReader) Read() (record map[string]string, err error) {
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
