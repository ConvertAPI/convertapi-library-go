package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/lib"
	"github.com/ConvertAPI/convertapi-go/param"
	"os"
)

func main() {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting PDF to JPG and compressing result files with ZIP")

	jpgRes := convertapi.ConvDef("docx", "jpg", param.NewPath("file", "assets/test.docx", nil))

	zipRes := convertapi.ConvDef("jpg", "zip", param.NewResult("files", jpgRes, nil))

	if cost, err := jpgRes.Cost(); lib.PrintErr(err) {
		fmt.Println("DOCX -> JPG conversion cost: ", cost)
	}

	if files, err := jpgRes.Files(); lib.PrintErr(err) {
		fmt.Println("DOCX -> JPG conversion result file count: ", len(files))
	}

	if cost, err := zipRes.Cost(); lib.PrintErr(err) {
		fmt.Println("JPG -> ZIP conversion cost: ", cost)
	}

	if files, errs := zipRes.ToPath("/tmp/result.zip"); errs == nil {
		fmt.Println("ZIP file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
