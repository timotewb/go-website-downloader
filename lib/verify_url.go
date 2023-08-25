package lib

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func VerifyURL(url string) (int, error) {
	fmt.Println(url)
	var client = http.Client{
		Transport: &http.Transport{
		  Dial: net.Dialer{Timeout: 2 * time.Second}.Dial,
		},
	  }
    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
       return 0, err
    }
    resp, err := client.Do(req)
    if err != nil {
       return 0, err
    }
    resp.Body.Close()
    return resp.StatusCode, nil
}