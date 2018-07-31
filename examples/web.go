package main

import (
	"../pkg/convertapi"
	"fmt"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.Convert("web", "pdf", []*convertapi.Param{
		convertapi.NewStringParam("url", "https://en.wikipedia.org/wiki/Data_conversion"),
		convertapi.NewStringParam("filename", "web-example"),
	}, nil)

	if files, errs := pptxRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
