package main

import (
	"flag"
	"fmt"
	"gowrapper"
)

var path string

func init() {
	flag.StringVar(&path, "i", "/dev/null", "input file path")
}

func main() {
	flag.Parse()
	i := 10
	f, err := gowrapper.Open(path)
	fmt.Printf("path = %s\n", path)
	gowrapper.Check(err)
	size := 5
	buffer := gowrapper.Make(size)
	n, err := gowrapper.Read(f, buffer, size)
	gowrapper.Check(err)
	fmt.Printf("%d bytes: %d\n", n, buffer[0])
	if buffer[0] == 35 {
		i = i + 1
	} else {
		i = i * 2
	}
	fmt.Printf("i = %d\n", i)
}
