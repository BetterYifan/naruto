package io

import (
	"bytes"
	"os"
	"testing"
)

func TestWriteReader(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("nihao")

	// 只读方式打开文件
	file, err := os.Open("file.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = b.ReadFrom(file)
	t.Log(string(b.Bytes()))

	b.String()

	b.Reset()

	b.WriteString("nidassasa")
	file, err = os.OpenFile("file.txt", os.O_RDWR, 0)
	b.WriteTo(file)
}
