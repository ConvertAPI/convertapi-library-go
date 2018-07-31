package convertapi

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func requestDelete(url string, client *http.Client) (err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err == nil {
		_, err = respExtractErr(client.Do(req))
	}
	return
}

func respExtractErr(r *http.Response, e error) (resp *http.Response, err error) {
	resp = r
	err = e
	if err == nil && resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		err = errors.New(buf.String())
	}
	return
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func pathExt(path string) string {
	return strings.Replace(filepath.Ext(path), ".", "", -1)
}

func addErr(errs *[]error, err error) bool {
	if err == nil {
		return true
	} else {
		*errs = append(*errs, err)
		return false
	}
}

func isDir(path string) bool {
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
