package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stdout, buffer)
	fmt.Fprintln(mw, "hello multiwriter....")
	fmt.Println(buffer)
}
