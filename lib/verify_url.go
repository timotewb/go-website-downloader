package lib

import (
	"net/http"
	"net/url"
	"time"
)

type ResponseType struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func ValidateURL(urlString string) ResponseType {
	time.Sleep(1 * time.Second)
	// create response object
	var r ResponseType
	// check url is valid
	_, err := url.ParseRequestURI(urlString)
	// error check
	if err != nil {
		r.Code = 1
		r.Message = err.Error()
		return r
	}
	// time.Sleep(1 * time.Second)
	r.Code = 0
	r.Message = ""
	return r
}

func VerifyURL(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
	   return 0, err
	}
	defer response.Body.Close()
	return response.StatusCode, nil
}