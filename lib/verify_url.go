package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
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
		rgx, _ := regexp.Compile(`favicon[^/]+.(ico|png)`)
		if StringInSlice(n.Data, l...) {
			for _, a := range n.Attr {
				if strings.Contains(a.Val, "favicon.ico"){
					fmt.Println("01")
					// this code is replicated - fix this
					if a.Val == "favicon.ico" {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + "/" + a.Val
						return
					} else if strings.HasPrefix(a.Val, "//") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
						return
					} else {
						fmt.Println(" --- Build new rule for GetFavicon() 01.")
					}
				} else if strings.Contains(a.Val, "favicon.png"){
					fmt.Println("02")
					if a.Val == "favicon.png" {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + "/" + a.Val
						return
					} else if strings.HasPrefix(a.Val, "//") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
						return
					} else {
						fmt.Println(" --- Build new rule for GetFavicon() 02.")
					}
				} else if rgx.MatchString(a.Val){
					fmt.Println("03")
					// this code is slightly different - Yay!
					if strings.HasPrefix(a.Val, "//") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "/") {
						r.FaviconURL = strings.TrimSuffix(u.String(),"/") + a.Val
						return
					} else if strings.HasPrefix(a.Val, "http") {
						r.FaviconURL = a.Val
						return
					} else {
						fmt.Println(" --- Build new rule for GetFavicon() 03.")
					}
				}
			}
		}
        // traverses the HTML of the webpage from the first child node
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            favi(c)
        }
	}
	favi(doc)
	return r, err
}