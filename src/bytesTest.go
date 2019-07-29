package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello world\n"))
	b.WriteTo(os.Stdout)

	fmt.Println(string(bytes.Replace([]byte("ooooink oink oink"), []byte("ooo"), []byte("xx"), 2)))
	fmt.Println(bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte("")))

	fmt.Println(bytes.LastIndexAny([]byte("aasdcdfggte&"), "a&e"))
	fmt.Println(bytes.LastIndexFunc([]byte("我asdff!2324!"), unicode.IsLetter))
	fmt.Println(string(bytes.Replace([]byte("asdaqweart qwe!"), []byte(""), []byte("x"), -1)))
}
