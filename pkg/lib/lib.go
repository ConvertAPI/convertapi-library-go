package lib

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RequestDelete(url string, client *http.Client) (err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err == nil {
		_, err = RespExtractErr(client.Do(req))
	}
	return
}

func RespExtractErr(r *http.Response, e error) (resp *http.Response, err error) {
	resp = r
	err = e
	if err == nil && resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		if _, err = buf.ReadFrom(resp.Body); err == nil {
			err = errors.New(buf.String())
		}
	}
	return
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func PathExt(path string) string {
	return strings.Replace(filepath.Ext(path), ".", "", -1)
}

func AddErr(errs *[]error, err error) bool {
	if err != nil {
		*errs = append(*errs, err)
		return false
	}
	return true
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func PrintErr(err error) bool {
	if err == nil {
		return true
	} else {
		fmt.Println("Error: ", err)
		return false
	}
}
