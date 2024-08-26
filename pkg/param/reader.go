package param

import (
	"bytes"
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"io"
	"net/url"
	"strings"
	"sync"
)

type ParamReader struct {
	Param
	reader   io.Reader
	fileName string
	config   *config.Config
	sync.Mutex
}

func NewReader(name string, reader io.Reader, filename string, conf *config.Config) *ParamReader {
	if conf == nil {
		conf = config.Default
	}
	return &ParamReader{*New(name), reader, filename, conf, sync.Mutex{}}
}

func (this *ParamReader) Prepare() error {
	this.Lock()
	defer this.Unlock()

	if this.values == nil {
		if err := this.Param.Prepare(); err != nil {
			return err
		}

		query := url.Values{}
		query.Add("filename", this.fileName)

		pathURL, err := url.Parse("/upload?" + query.Encode())
		if err != nil {
			return err
		}

		uploadURL := this.config.BaseURL.ResolveReference(pathURL)
		resp, err := this.config.HttpClient.Post(uploadURL.String(), "application/octet-stream", this.reader)

		if err != nil {
			return err
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return err
		}
		this.values = []string{buf.String()}
	}
	return nil
}

func (this *ParamReader) Values() ([]string, error) {
	err := this.Prepare()
	return this.values, err
}

func (this *ParamReader) String() string {
	return fmt.Sprintf("%s: %s -> %s", this.name, this.fileName, strings.Join(this.values, " "))
}
