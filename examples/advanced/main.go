package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
	"net/http"
	"net/url"
	"os"
)

func main() {
	token := os.Getenv("API_TOKEN")

	// Using convertapi.com server in Europe
	u, _ := url.Parse("https://eu-api.convertapi.io")

	// Using HTTP proxy server
	proxy, _ := url.Parse("http://127.0.0.1:8888")
	transport := &http.Transport{Proxy: http.ProxyURL(proxy)}

	// Setting this configuration as default
	config.Default = config.New(token, u, transport)

	fmt.Println("Converting remote PPTX to PDF")
	fileParam := param.NewString("file", "https://cdn.convertapi.com/cara/testfiles/presentation.pptx")
	pptxRes := convertapi.Convert("pptx", "pdf", []param.Parameter{fileParam}, nil)

	if files, errs := pptxRes.ToPath("/tmp/converted.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
		fmt.Println("Deleting source file from convertapi.com server")
		fileParam.Delete(nil)
		fmt.Println("Deleting result files from convertapi.com server")
		pptxRes.Delete()
	} else {
		fmt.Println(errs)
	}
}
