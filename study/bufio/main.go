package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	ReadSlice()
}

func ReadSlice() {
	reader := bufio.NewReaderSize(strings.NewReader("123456789\nabcdefg\nhijklmn"), 11)
	line, _ := reader.ReadSlice('\n')
	fmt.Println(string(line))
	bytes, _ := reader.ReadSlice('\n')
	fmt.Println(string(line))
	fmt.Println(string(bytes))
}
