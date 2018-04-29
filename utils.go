package goCsv

import (
	"encoding/csv"
	"io"
)

// CountLines 统计csv文件的记录数
func CountLines(r io.Reader) (n int, err error) {
	reader := csv.NewReader(r)

	// header
	if _, err = reader.Read(); err != nil {
		return n, err
	}

	for {
		_, err = reader.Read()
		if err == io.EOF {
			return n, nil
		} else if err != nil {
			return n, err
		}
		n++
	}
	return
}
