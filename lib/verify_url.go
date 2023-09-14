package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type ResponseType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Url string `json:"url"`
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

func GetFavicon(r ResponseType, page *http.Response) (ResponseType, error){

    doc, err := html.Parse(page.Body)
    if err != nil {
        fmt.Println("Error:", err)
		r.Code = 3
		r.Message = err.Error()
        return r, err
    }
	defer page.Body.Close()

	// base url
	u, err := url.Parse(r.Url)
	if err != nil {
        fmt.Println("Error:", err)
		r.Code = 3
		r.Message = err.Error()
        return r, err
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""

	// find favicon.ico
 	var favi func(*html.Node)
	favi = func(n *html.Node) {
		l := []string{"link", "meta"}
		if StringInSlice(n.Data, l...) {
			for _, a := range n.Attr {
				if strings.Contains(a.Val, "favicon.ico"){
					if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
					}
				} else if strings.Contains(a.Val, "favicon.png"){
					if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
					}
				} else if strings.Contains(a.Val, "favicon-32x32.png"){
					if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
					}
				} // add regex search for favicon*.ico and favicon*.png
			}
		}
        // traverses the HTML of the webpage from the first child node
        for c := n.FirstChild; c != nil; c = c.NextSibling {
			fmt.Println(c.Data)
            favi(c)
        }
	}
	favi(doc)
	return r, err
}