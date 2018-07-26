package convertapi

import (
	"net/http"
)

func requestDelete(url string, client *http.Client) {
	if req, err := http.NewRequest(url, http.MethodDelete, nil); err == nil {
		client.Do(req)
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
