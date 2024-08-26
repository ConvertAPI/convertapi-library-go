package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	fmt.Println("Creating PDF thumbnail")
	extractRes := convertapi.ConvDef("pdf", "jpg",
		param.NewPath("file", "assets/test.pdf", nil),
		param.NewString("pagerange", "1"),
	)

	jpgRes := convertapi.ConvDef("jpg", "jpg",
		param.NewResult("file", extractRes, nil),
		param.NewBool("scaleimage", true),
		param.NewBool("scaleproportions", true),
		param.NewInt("imageheight", 300),
	)

	if files, errs := jpgRes.ToPath("/tmp/thumbnail.jpg"); errs == nil {
		fmt.Println("Thumbnail saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}

}
