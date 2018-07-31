package main

import (
	"../pkg/convertapi"
	"fmt"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Creating PDF thumbnail")
	extractRes := convertapi.Convert("pdf", "extract", []*convertapi.Param{
		convertapi.NewFilePathParam("file", "assets/test.pdf", nil),
		convertapi.NewStringParam("pagerange", "1"),
	}, nil)

	jpgRes := convertapi.Convert("pdf", "jpg", []*convertapi.Param{
		convertapi.NewResultParam("file", extractRes, nil),
		convertapi.NewBoolParam("scaleimage", true),
		convertapi.NewBoolParam("scaleproportions", true),
		convertapi.NewIntParam("imageheight", 300),
	}, nil)

	if files, errs := jpgRes.ToPath("/tmp/thumbnail.jpg"); errs == nil {
		fmt.Println("Thumbnail saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
