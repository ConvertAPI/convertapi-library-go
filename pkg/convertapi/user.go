package convertapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type User struct {
	Active      bool
	ApiKey      int
	Email       string
	FullName    string
	SecondsLeft int
	Secret      string
	Status      string
}

func UserInfo(config *Config) (user *User, err error) {
	if config == nil {
		config = Default
	}
	query := url.Values{}
	query.Add("secret", config.Secret)
	path := fmt.Sprintf("/user?%s", query.Encode())
	pathURL, err := url.Parse(path)
	if err != nil {
		return
	}
	userURL := config.BaseURL.ResolveReference(pathURL)
	resp, err := respExtractErr(config.HttpClient.Get(userURL.String()))
	if err == nil {
		user = &User{}
		err = json.NewDecoder(resp.Body).Decode(user)
	}

	return
}
