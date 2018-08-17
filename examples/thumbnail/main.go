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

	fmt.Println("Creating PDF thumbnail")
	extractRes := convertapi.Convert("pdf", "extract", []param.IParam{
		param.NewPath("file", "assets/test.pdf", nil),
		param.NewString("pagerange", "1"),
	}, nil)

	jpgRes := convertapi.Convert("pdf", "jpg", []param.IParam{
		param.NewResult("file", extractRes, nil),
		param.NewBool("scaleimage", true),
		param.NewBool("scaleproportions", true),
		param.NewInt("imageheight", 300),
	}, nil)

	if files, errs := jpgRes.ToPath("/tmp/thumbnail.jpg"); errs == nil {
		fmt.Println("Thumbnail saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
