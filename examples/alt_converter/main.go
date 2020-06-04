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

	fmt.Println("Converting DOCX to PDF using OpenOffice alternative converter")

	pdfRes := convertapi.ConvDef("docx", "pdf",
		param.NewPath("file", "assets/test.docx", nil),
		param.NewString("converter", "openoffice"))

	if files, errs := pdfRes.ToPath("/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
