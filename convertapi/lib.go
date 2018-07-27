package convertapi

import (
	"net/http"
)

func requestDelete(url string, client *http.Client) (err error) {
	if req, err := http.NewRequest(url, http.MethodDelete, nil); err == nil {
		_, err = client.Do(req)
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
