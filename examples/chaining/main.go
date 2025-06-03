package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/lib"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("API_TOKEN")) // Get your token at https://www.convertapi.com/a/authentication

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