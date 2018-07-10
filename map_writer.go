package goCsv

import (
	"encoding/csv"
	"errors"
	"io"
)

type MapWriter struct {
	Writer         *csv.Writer
	fieldnames     []string
	fieldNotSetErr bool // 字段未设置时，是否报错
}

func NewMapWriter(w io.Writer, fieldnames []string) *MapWriter {
	return &MapWriter{
		Writer:     csv.NewWriter(w),
		fieldnames: fieldnames,
	}
}

// NewMapWriterSimple 简化的writer对象，可以通过SetHeader方法来设置Header
func NewMapWriterSimple(w io.Writer) *MapWriter {
	return &MapWriter{
		Writer: csv.NewWriter(w),
	}
}

func (w *MapWriter) SetHeader(fieldnames []string) {
	w.fieldnames = fieldnames
}

// SetFieldNotSetErr 字段未设置时，是否报错
// 默认为false，即当字段为设置时，会自动使用空字符串补充
func (w *MapWriter) SetFieldNotSetErr(fieldNotSetErr bool) {
	w.fieldNotSetErr = fieldNotSetErr
}

func (w *MapWriter) WriteHeader() (err error) {
	return w.Writer.Write(w.fieldnames)
}

// Init 配置csv的基本参数
func (w *MapWriter) Init(params csv.Writer) {
	w.Writer.Comma = params.Comma
}

func (w *MapWriter) WriteRow(row map[string]string) (err error) {
	var ok bool
	var val string
	var record = make([]string, len(w.fieldnames))
	record = record[0:0]
	for _, i := range w.fieldnames {
		if val, ok = row[i]; ok {
			record = append(record, val)
		} else {
			if w.fieldNotSetErr {
				return errors.New("the field name is not exist: " + i)
			} else {
				record = append(record, "")
			}
		}
	}

	return w.Writer.Write(record)
}

func (w *MapWriter) WriteRows(rows []map[string]string) (err error) {
	for _, row := range rows {
		if err = w.WriteRow(row); err != nil {
			return err
		}
	}
	return nil
}

// Flush 将数据刷到磁盘
func (w *MapWriter) Flush() {
	w.Writer.Flush()
}
