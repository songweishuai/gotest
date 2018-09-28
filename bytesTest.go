package main

import (
	"bytes"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello world\n"))
	b.WriteTo(os.Stdout)
}
