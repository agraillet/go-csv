package goCsv

import (
	"encoding/csv"
	"errors"
	"io"
)

type MapWriter struct {
	Writer     *csv.Writer
	fieldnames []string
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
			return errors.New("the field name is not exist: " + i)
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
