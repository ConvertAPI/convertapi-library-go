package main

import (
	"fmt"
	convertapi "github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting WEB page to PDF")
	webRes := convertapi.ConvDef("web", "pdf",
		param.NewString("url", "https://en.wikipedia.org/wiki/Data_conversion"),
		param.NewString("filename", "web-example"),
	)

	if files, errs := webRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
