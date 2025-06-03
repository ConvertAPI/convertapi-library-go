package main

import (
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
)

// Example of DOCX to PDF conversion using vanilla go.
// File is converted without temporary storing files on ConverAPI servers.
// No error handling to make an example easier to read.
// No file buffering to memory. To store file to memory in general is a bad idea!
func main() {
	pipeReader, pipeWriter := io.Pipe()
	multipartWriter := multipart.NewWriter(pipeWriter)

	// Adjust URL according your converter and set API_TOKEN.
	req, _ := http.NewRequest("POST", "https://v2.convertapi.com/convert/docx/to/pdf?secret=API_TOKEN", pipeReader)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Accept", "multipart/mixed")

	go func(mpWriter *multipart.Writer) {
		defer pipeWriter.Close()
		defer multipartWriter.Close()
		filePart, _ := mpWriter.CreateFormFile("file", "test.docx")
		f, _ := os.Open("assets/test.docx") // Open file that needs to be converted
		defer f.Close()
		f.WriteTo(filePart)
		// If conversion takes multiple source files, it can be added here, as an example above.
	}(multipartWriter)

	resp, _ := http.DefaultClient.Do(req)
	save(resp, "/tmp/result.pdf")
}

func save(resp *http.Response, file string) {
	defer resp.Body.Close()

	mediaType := resp.Header.Get("Content-Type")
	_, params, _ := mime.ParseMediaType(mediaType)
	multipartReader := multipart.NewReader(resp.Body, params["boundary"])

	part, _ := multipartReader.NextPart()
	resFile, _ := os.Create(file)
	defer resFile.Close()
	io.Copy(resFile, part)
	// If conversion returns multiple files, it can be red using multipartReader.NextPart() multiple times.
}