package main

import (
	"../pkg/convertapi"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	secret := os.Getenv("CONVERTAPI_SECRET")

	// Using convertapi.com server in Europe
	domain, _ := url.Parse("https://eu.v2.convertapi.com")

	// Using HTTP proxy server
	proxy, _ := url.Parse("http://127.0.0.1:8888")
	transport := &http.Transport{Proxy: http.ProxyURL(proxy)}

	// Setting this configuration as default
	convertapi.Default = convertapi.NewConfig(secret, domain, transport)

	fmt.Println("Converting remote PPTX to PDF")
	fileParam := convertapi.NewStringParam("file", "https://cdn.convertapi.com/cara/testfiles/presentation.pptx")
	pptxRes := convertapi.Convert("pptx", "pdf", []*convertapi.Param{fileParam}, nil)

	if files, errs := pptxRes.ToPath("/tmp/converted.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
		fmt.Println("Deleting source file from convertapi.com server")
		fileParam.Delete()
		fmt.Println("Deleting result files from convertapi.com server")
		pptxRes.Delete()
	} else {
		fmt.Println(errs)
	}
}
