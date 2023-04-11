package main

import (
	"flag"
	"fmt"
)

func main() {
	age := flag.Int("age", 18, "")
	flag.Parse()
	fmt.Println(age)
}
