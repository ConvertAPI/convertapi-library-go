package main

import (
	"../convertapi"
	"fmt"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	if file, errs := convertapi.ConvertPath("test-files/test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
