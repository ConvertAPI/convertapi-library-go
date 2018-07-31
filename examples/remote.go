package main

import (
	"../convertapi"
	"fmt"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.Convert("pptx", "pdf", []*convertapi.Param{
		convertapi.NewStringParam("file", "https://cdn.convertapi.com/cara/testfiles/presentation.pptx"),
	}, nil)

	if files, errs := pptxRes.ToPath("/tmp/converted.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
