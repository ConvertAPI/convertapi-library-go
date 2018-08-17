package convertapi

import (
	"github.com/ConvertAPI/convertapi-go/lib"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ResFile struct {
	client   *http.Client
	resp     *http.Response
	FileName string
	FileSize int
	Url      string
}

func (this *ResFile) download() (err error) {
	if this.resp == nil {
		this.resp, err = this.client.Get(this.Url)
	}
	return
}

func (this *ResFile) Read(p []byte) (n int, err error) {
	err = this.download()
	if err == nil {
		n, err = this.resp.Body.Read(p)
		if err != nil {
			this.resp.Body.Close()
			this.resp = nil
		}
	}
	return
}

func (this *ResFile) ToFile(file *os.File) (err error) {
	_, err = io.Copy(file, this)
	return
}

func (this *ResFile) ToPath(path string) (file *os.File, err error) {
	if info, e := os.Stat(path); e == nil && info.IsDir() {
		path = filepath.Join(path, this.FileName)
	}

	if file, err = os.Create(path); err == nil {
		defer file.Close()
		err = this.ToFile(file)
	}
	return
}

func (this *ResFile) Delete() error {
	return lib.RequestDelete(this.Url, this.client)
}
