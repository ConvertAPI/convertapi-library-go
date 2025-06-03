# ConvertAPI Go Client

![example workflow](https://github.com/ConvertAPI/convertapi-go/actions/workflows/tests.yml/badge.svg)

## Convert your files with our online file conversion API

ConvertAPI helps in converting various file formats. Creating PDF and Images from various sources like Word, Excel, Powerpoint, images, web pages or raw HTML codes. Merge, Encrypt, Split, Repair and Decrypt PDF files and many other file manipulations. You can integrate it into your application in just a few minutes and use it easily.

The ConvertAPI-Go library makes it easier to use the Convert API from your Go projects without having to build your own API calls.
You can get your free API token at https://www.convertapi.com/a/authentication

## Installation

Execute this command in your GOPATH:

```shell
go get github.com/ConvertAPI/convertapi-go
```

## Usage

### Configuration

You can get your free API token at https://www.convertapi.com/a/authentication

```go
config.Default = config.NewDefault(os.Getenv("api-token"))
```

### File conversion

Convert DOCX file to PDF example.
All supported formats and options can be found [here](https://www.convertapi.com/conversions).

```go
pdfRes := convertapi.ConvDef("docx", "pdf", param.NewPath("file", "test.docx", nil))

// save to file
pdfRes.ToPath("/tmp/result.pdf")
```

Other result operations:

```go
// save all result files to folder
res.ToPath("/tmp")

// get second file in result
files, err := res.Files()
secondFile := files[1]

// result implements Reader interface so it is possible to stream result
io.Copy(myReader, res)

// get conversion cost
cost, err := res.Cost()
```

#### Convert remote file

```go
pptxRes := convertapi.ConvDef("pptx", "pdf", 
    param.NewString("file", "https://cdn.convertapi.com/cara/testfiles/presentation.pptx"),
)
```

#### Additional conversion parameters

ConvertAPI accepts extra conversion parameters depending on selected file formats.
All conversion parameters and explanations can be found [here](https://www.convertapi.com/conversions).

```go
jpgRes := convertapi.ConvDef("pdf", "jpg",
    param.NewResult("file", extractRes, nil),
    param.NewBool("scaleimage", true),
    param.NewBool("scaleproportions", true),
    param.NewInt("imageheight", 300),
)
```

### User information

You can always check your remaining seconds programmatically by fetching [user information](https://www.convertapi.com/doc/user).

```go
user, err := convertapi.UserInfo(nil)
conversionsTotal := user.ConversionsTotal
conversionsConsumed := user.ConversionsConsumed
```

### Alternative domain

How to set an alternative domain can be found in the [advanced example](https://github.com/ConvertAPI/convertapi-go/blob/master/examples/advanced/main.go).
Dedicated to the region [domain list](https://www.convertapi.com/doc/servers-location).

### More examples

Find more advanced examples in the [examples](https://github.com/ConvertAPI/convertapi-go/tree/master/examples) folder.

#### Converting your first file, full example:

ConvertAPI is designed to make converting file super easy. The following snippet shows how easy it is to get started. Let's convert a WORD DOCX file to PDF:

```go
package main

import (
	"github.com/ConvertAPI/convertapi-go"
	"fmt"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("api-token"))

	if file, errs := convertapi.ConvertPath("test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
```

This is the bare-minimum to convert a file using the ConvertAPI client, but you can do a great deal more with the ConvertAPI Go library.
Take special note that you should replace `api-token` with the token you obtained in item two of the pre-requisites.


### Conversion Without Temporarily Storing Files on the ConvertAPI File Server
By design, the ConvertAPI library uses a temporary file server to store files.
A temporary file server is useful for conversion chaining and generally improves performance.
However, the [Vanilla Go Example](examples/direct/main.go) demonstrates how to efficiently convert files without using the ConvertAPI file server.

In this example, both the request and response use the `multipart` content type. This approach allows you to convert multiple files or return conversion results as multiple files.
Using `multipart` is much more efficient compared to JSON with base64 encoding.


### Issues &amp; Comments
Please leave all comments, bugs, requests, and issues on the Issues page. We'll respond to your request ASAP!

### License
The ConvertAPI Go Library is licensed under the [MIT](http://www.opensource.org/licenses/mit-license.php "Read more about the MIT license form") license.
Refer to the [LICENSE](https://github.com/ConvertAPI/convertapi-go/blob/master/LICENSE) file for more information.