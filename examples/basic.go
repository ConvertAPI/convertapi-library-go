package main

import (
	"../convertapi"
	"os"
)

func main() {
	convertapi.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	res := convertapi.Convert("pdf", "jpg", []*convertapi.Param{
		convertapi.NewFilePathParam("file", "/tmp/test.pdf", nil),
	}, nil)

	//fmt.Print(res.Response())
	res.ToPath("/tmp/")
	res.Delete()
}
