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