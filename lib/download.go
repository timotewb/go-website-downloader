package lib

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetFavicon(url string) (string, error) {
	Logger(url)
	// Send an HTTP GET request to the example.com web page
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return "", err
    }
    defer resp.Body.Close()

    // Use the html package to parse the response body from the request
    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return "", err
    }

	// find favicon.ico
	var faviconURL string
	var favi func(*html.Node)
	favi = func(n *html.Node) {
		if n.Data == "link" {
			for _, a := range n.Attr {
				if strings.Contains(a.Val, "favicon.ico"){
					if strings.HasPrefix(a.Val, "/") {
						faviconURL = strings.TrimSuffix(url,"/") + a.Val
					} else if strings.HasPrefix(a.Val, "http") {
						faviconURL = a.Val
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
	Logger(faviconURL)
	Logger("Done")
	return faviconURL, nil
}