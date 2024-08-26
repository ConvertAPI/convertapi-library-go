package main

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"os"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	if user, err := convertapi.UserInfo(nil); err == nil {
		fmt.Println("User information: ")
		fmt.Printf("%+v\n", user)
	} else {
		fmt.Println(err)
	}
}
