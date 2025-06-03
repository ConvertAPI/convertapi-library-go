package main

import (
	"fmt"
	convertapi "github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("API_TOKEN")) // Get your token at https://www.convertapi.com/a/authentication

	fmt.Println("Converting HTML from the stream to TXT")

	reader := strings.NewReader("<!DOCTYPE html><html><head><title>My First Heading</title></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>\n")

	htmlRes := convertapi.ConvDef("html", "txt",
		// Reading source file body from the io.Reader
		param.NewReader("file", reader, "page.html", nil),
	)

	fmt.Println("TXT file content:")

	// Copy the content of the reader to stdout.
	if _, err := io.Copy(os.Stdout, htmlRes); err != nil {
		log.Fatalf("failed to copy data: %s", err)
	}
}