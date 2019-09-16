package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	header := strings.NewReader("<msg>")
	body := strings.NewReader("Hello Multi Reader")
	footer := strings.NewReader("</msg>")

	//  for _, r := range []io.Reader{header, body, footer} {
	// 	_, err := io.Copy(os.Stdout, r)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }

	r := io.MultiReader(header, body, footer)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}

}
