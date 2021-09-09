package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"os"
)

func main() {
	config.Default.Token = os.Getenv("CONVERTAPI_TOKEN")   // Get your token at https://www.convertapi.com/a/auth
	config.Default.ApiKey = os.Getenv("CONVERTAPI_APIKEY") // Get your API key at https://www.convertapi.com/a/auth

	if file, errs := convertapi.ConvertPath("assets/test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
