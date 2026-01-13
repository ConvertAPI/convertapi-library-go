package convertapi

import (
	"github.com/ConvertAPI/convertapi-go/pkg/lib"
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
	FileID   string
	URL      string
}

func (f *ResFile) download() (err error) {
	if f.resp == nil {
		f.resp, err = f.client.Get(f.URL)
	}
	return
}

func (f *ResFile) Read(p []byte) (n int, err error) {
	err = f.download()
	if err == nil {
		n, err = f.resp.Body.Read(p)
		if err != nil {
			f.resp.Body.Close()
			f.resp = nil
		}
	}
	return
}

func (f *ResFile) ToFile(file *os.File) (err error) {
	_, err = io.Copy(file, f)
	return
}

func (f *ResFile) ToPath(path string) (file *os.File, err error) {
	if info, e := os.Stat(path); e == nil && info.IsDir() {
		path = filepath.Join(path, f.FileName)
	}

	if file, err = os.Create(path); err == nil {
		defer file.Close()
		err = f.ToFile(file)
	}
	return
}

func (f *ResFile) Delete() error {
	return lib.RequestDelete(f.URL, f.client)
}
