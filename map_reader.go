// 参考Python的DictReader
package goCsv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type MapReader struct {
	Reader *csv.Reader

	skip       int  // 跳过前面若干条记录
	limit      int  // 只读取若干条有效的记录
	isLimit    bool // 是否限制返回记录数量
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

// SetSkip 设置跳过前面的若干条记录
func (r *MapReader) SetSkip(skip int) {
	r.skip = skip
}

// SetLimit 设置只提取若干条记录
func (r *MapReader) SetLimit(limit int) {
	r.limit = limit
	r.isLimit = true
}

// SetFieldnames 指定csv文件的字段名
// 如果不指定的话，则默认使用csv文件的第一行作为字段名
func (r *MapReader) SetFieldnames(fieldnames []string) {
	r.fieldnames = fieldnames
}

// GetFieldnames 获取csv文件的header
// csv文件在处理的时候，可能会最后会出现空字段
// Error: Multiple indices with the same name ''
func (r *MapReader) GetFieldnames() (fieldnames []string, err error) {
	if len(r.fieldnames) == 0 {
		// 如果没有设置字段名，则默认为csv的第一行为字段名
		if fieldnames, err = r.Reader.Read(); err != nil {
			return nil, err
		}
	}

	// 格式化fieldnames
	emptyCnt := 0
	for i := len(fieldnames) - 1; i >= 0; i-- {
		if fieldnames[i] == "" {
			emptyCnt++
		} else {
			break
		}
	}
	fieldnames = fieldnames[:-emptyCnt]
	r.fieldnames = fieldnames
	return fieldnames, nil
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

	for {
		if r.skip <= 0 {
			break
		}
		// 跳过前面的若干记录
		if _, err = r.Reader.Read(); err != nil {
			return nil, err
		}
		r.skip--
	}

	if r.isLimit {
		if r.limit == 0 {
			// 已经读完所有记录
			return nil, io.EOF
		} else {
			r.limit--
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
