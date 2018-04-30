// 对原reader的封装
package goCsv

import (
	"encoding/csv"
	"io"
)

type Reader struct {
	Reader *csv.Reader

	skip       int  // 跳过前面若干条记录
	limit      int  // 只读取若干条有效的记录
	isLimit    bool // 是否限制返回记录数量
	fieldnames []string
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		Reader: csv.NewReader(r),
	}
}

// Init 配置csv的基本参数
func (r *Reader) Init(params csv.Reader) {
	r.Reader.Comma = params.Comma
	r.Reader.LazyQuotes = params.LazyQuotes
	r.Reader.TrimLeadingSpace = params.TrimLeadingSpace
}

// SetSkip 设置跳过前面的若干条记录
func (r *Reader) SetSkip(skip int) {
	r.skip = skip
}

// SetLimit 设置只提取若干条记录
func (r *Reader) SetLimit(limit int) {
	r.limit = limit
	r.isLimit = true
}

func (r *Reader) GetFieldnames() (fieldnames []string, err error) {
	if len(r.fieldnames) == 0 {
		// 如果没有设置字段名，则默认为csv的第一行为字段名
		if r.fieldnames, err = r.Reader.Read(); err != nil {
			return nil, err
		}
	}
	return r.fieldnames, nil
}

// Read 读取一行记录
func (r *Reader) Read() (record []string, err error) {
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

	return r.Reader.Read()
}

// ReadAll 读取全部的内容
func (r *Reader) ReadAll() (records [][]string, err error) {
	var record []string
	for record, err = r.Read(); err == nil; record, err = r.Read() {
		records = append(records, record)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return records, nil
}
