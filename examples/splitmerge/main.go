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

	fmt.Println("Creating PDF with the first and the last pages")
	splitRes := convertapi.ConvDef("pdf", "split",
		param.NewPath("file", "assets/test.pdf", nil),
	)

	mergeRes := convertapi.ConvDef("pdf", "merge",
		param.NewResultIdx("files", splitRes, 0, nil),
		param.NewResultIdx("files", splitRes, -1, nil),
	)

	if files, errs := mergeRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
