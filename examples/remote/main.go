package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("API_TOKEN")) // Get your token at https://www.convertapi.com/a/authentication

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.ConvDef("pptx", "pdf",
		param.NewString("file", "https://cdn.convertapi.com/public/files/demo.pptx"))

	if files, errs := pptxRes.ToPath("/tmp/converted.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}