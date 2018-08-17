package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"os"
)

func main() {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting DOCX to PDF and JPG in parallel using same source file")
	fileParam := param.NewPath("file", "assets/test.docx", nil)
	pdfRes := convertapi.Convert("docx", "pdf", []param.IParam{fileParam}, nil)
	jpgRes := convertapi.Convert("docx", "jpg", []param.IParam{fileParam}, nil)

	// Downloading and saving files also in parallel
	c1 := save(pdfRes)
	c2 := save(jpgRes)
	<-c1
	<-c2
}

func save(res *convertapi.Result) (finish chan struct{}) {
	finish = make(chan struct{})
	go func() {
		defer close(finish)
		if files, errs := res.ToPath("/tmp"); errs == nil {
			fmt.Println("File saved to: ", files[0].Name())
		} else {
			fmt.Println(errs)
		}
	}()
	return
}
