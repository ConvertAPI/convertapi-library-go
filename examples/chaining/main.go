package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting PDF to JPG and compressing result files with ZIP")

	jpgRes := convertapi.Convert("docx", "jpg", []*convertapi.Param{
		convertapi.NewFilePathParam("file", "assets/test.docx", nil),
	}, nil)

	zipRes := convertapi.Convert("jpg", "zip", []*convertapi.Param{
		convertapi.NewResultParam("files", jpgRes, nil),
	}, nil)

	if cost, err := jpgRes.Cost(); convertapi.PrintErr(err) {
		fmt.Println("DOCX -> JPG conversion cost: ", cost)
	}

	if files, err := jpgRes.Files(); convertapi.PrintErr(err) {
		fmt.Println("DOCX -> JPG conversion result file count: ", len(files))
	}

	if cost, err := zipRes.Cost(); convertapi.PrintErr(err) {
		fmt.Println("JPG -> ZIP conversion cost: ", cost)
	}

	if files, errs := zipRes.ToPath("/tmp/result.zip"); errs == nil {
		fmt.Println("ZIP file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
