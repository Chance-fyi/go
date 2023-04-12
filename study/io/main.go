package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	//NewSectionReader()
	//pipe()
	//multiReader()
}

func NewSectionReader() {
	reader := bytes.NewReader([]byte("0123456789"))
	sr := io.NewSectionReader(reader, 4, 6)
	b := make([]byte, 10)
	sr.Read(b)
	fmt.Println(string(b))
}

func pipe() {
	reader, writer := io.Pipe()
	go func(r *io.PipeReader) {
		buf := make([]byte, 20)
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("reader: ", n, string(buf))
	}(reader)
	go func(r *io.PipeWriter) {
		data := []byte("hello world")
		fmt.Println(string(data))
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("writer: ", n, string(data))
	}(writer)
	time.Sleep(time.Second * 3)
}

func multiReader() {
	mr1 := io.MultiReader(strings.NewReader("def"), strings.NewReader("ghi"))
	mr2 := io.MultiReader(strings.NewReader("abc"), mr1)

	buf := make([]byte, 9)
	io.ReadFull(mr2, buf)
	fmt.Println(string(buf))
}
