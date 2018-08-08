# ConvertAPI Go Client

[![Build Status](https://secure.travis-ci.org/ConvertAPI/convertapi-go.svg)](http://travis-ci.org/ConvertAPI/convertapi-go)

## Convert your files with our online file conversion API

The ConvertAPI helps converting various file formats.
Creating PDF and Images from various sources like Word, Excel, Powerpoint, images, web pages or raw HTML codes.
Merge, Encrypt, Split, Repair and Decrypt PDF files.
And many others files manipulations.
In just few minutes you can integrate it into your application and use it easily.

The ConvertAPI-Go library makes it easier to use the Convert API from your Go projects without having to build your own API calls.
You can get your free API secret at https://www.convertapi.com/a

## Installation

Execute in your GOPATH this command:

```shell
go get github.com/ConvertAPI/convertapi-go
```

## Usage

### Configuration

You can get your secret at https://www.convertapi.com/a

```go
convertapi.Default.Secret = "YOUR API SECRET"
```

### File conversion

Example to convert DOCX file to PDF.
All supported formats and options can be found [here](https://www.convertapi.com).

```go
pdfRes := convertapi.Convert("docx", "pdf", []*convertapi.Param{
    convertapi.NewFilePathParam("file", "test.docx", nil),
}, nil)

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
pptxRes := convertapi.Convert("pptx", "pdf", []*convertapi.Param{
    convertapi.NewStringParam("file", "https://cdn.convertapi.com/cara/testfiles/presentation.pptx"),
}, nil)
```

#### Additional conversion parameters

ConvertAPI accepts extra conversion parameters depending on converted formats.
All conversion parameters and explanations can be found [here](https://www.convertapi.com).

```go
jpgRes := convertapi.Convert("pdf", "jpg", []*convertapi.Param{
    convertapi.NewResultParam("file", extractRes, nil),
    convertapi.NewBoolParam("scaleimage", true),
    convertapi.NewBoolParam("scaleproportions", true),
    convertapi.NewIntParam("imageheight", 300),
}, nil)
```

### User information

You can always check remaining seconds amount by fetching [user information](https://www.convertapi.com/doc/user).

```go
user, err := convertapi.UserInfo(nil)
secondsLeft := user.SecondsLeft
```

### More examples

You can find more advanced examples in the [examples](https://github.com/ConvertAPI/convertapi-go/tree/master/examples) folder.

#### Converting your first file, full example:

ConvertAPI is designed to make converting file super easy, the following snippet shows how easy it is to get started. Let's convert WORD DOCX file to PDF:

```go
package main

import (
	"github.com/ConvertAPI/convertapi-go"
	"fmt"
	"os"
)

func main() {
	convertapi.Default.Secret = "YOUR API SECRET"

	if file, errs := convertapi.ConvertPath("test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
```

This is the bare-minimum to convert a file using the ConvertAPI client, but you can do a great deal more with the ConvertAPI Go library.
Take special note that you should replace `YOUR API SECRET` with the secret you obtained in item two of the pre-requisites.

### Issues &amp; Comments
Please leave all comments, bugs, requests, and issues on the Issues page. We'll respond to your request ASAP!

### License
The ConvertAPI Go Library is licensed under the [MIT](http://www.opensource.org/licenses/mit-license.php "Read more about the MIT license form") license.
Refer to the [LICENSE](https://github.com/ConvertAPI/convertapi-go/blob/master/LICENSE) file for more information.