package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ValidateURL(urlString string) {
	u, err := url.ParseRequestURI(urlString)
	fmt.Println("Sleeping")
	time.Sleep(5 * time.Second)
	fmt.Println(u, err)
}

func VerifyURL(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
	   return 0, err
	}
	defer response.Body.Close()
	return response.StatusCode, nil
}