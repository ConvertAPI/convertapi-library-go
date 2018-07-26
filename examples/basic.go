package main

import (
	"../convertapi"
	"fmt"
	"os"
)

func main() {
	convertapi.Secret(os.Getenv("CONVERTAPI_SECRET"))
	res := convertapi.Convert("pdf", "jpg", []*convertapi.Param{
		convertapi.NewFilePathParam("file", "/tmp/test.pdf", nil),
	}, nil)

	fmt.Print(res.Response())
}
