package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ResponseType struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Url        string `json:"url"`
	FaviconURL string `json:"favicon_url"`
}

func ValidateURL(urlString string) (ResponseType, error) {
	time.Sleep(1 * time.Second)
	// create response object
	var r ResponseType
	// check url is valid
	u, err := url.ParseRequestURI(urlString)
	r.Url = u.String()
	fmt.Println("--- Parse Done")
	// error check
	if err != nil {
		r.Code = 1
		r.Message = err.Error()
		return r, err
	}
	fmt.Println("--- Error Check Done")
	r.Code = 0
	r.Message = "Success"
	fmt.Println("--- ready to return")
	return r, err
}

func VerifyURL(r ResponseType) (ResponseType, *http.Response, error) {
	resp, err := http.Get(r.Url)
	if err != nil {
		r.Code = 2
		r.Message = err.Error()
		return r, resp, err
	}
	return r, resp, nil
}

func GetFavicon(r ResponseType, page *http.Response) (ResponseType, error) {

	//----------------------------------------------------------------------------------------
	// https://icons.duckduckgo.com/ip3/
	//----------------------------------------------------------------------------------------
	r.FaviconURL = "https://icons.duckduckgo.com/ip3/" + strings.Split(strings.TrimPrefix(strings.TrimPrefix(r.Url, "http://"), "https://"), "/")[0] + ".ico"
	resp, err := http.Get(r.Url)
	if err == nil && resp.StatusCode == http.StatusOK {
		return r, nil
	}

	//----------------------------------------------------------------------------------------
	// https://www.icon.horse/icon/
	//----------------------------------------------------------------------------------------
	r.FaviconURL = "https://www.icon.horse/icon/" + strings.Split(strings.TrimPrefix(strings.TrimPrefix(r.Url, "http://"), "https://"), "/")[0]
	resp, err = http.Get(r.Url)
	if err == nil && resp.StatusCode == http.StatusOK {
		return r, nil
	}

	return r, nil
}
