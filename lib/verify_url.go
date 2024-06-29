package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
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
	fmt.Println(r.Url)
	page, err := http.Get(r.Url)
	if err != nil {
		return r, err
	}
	// Read the body of the response
	body, err := io.ReadAll(page.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a tokenizer
	tokenizer := html.NewTokenizer(bytes.NewReader(body))
	// define tags to search in
	tags := []string{"link", "meta"}

	// Iterate over the tokens
	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			// End of document reached
			return r, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			// fmt.Printf("Token Type: %v, Token Data: %v\n", tt, tokenizer.Token().Data)
			// attrs := tokenizer.Token().Attr
			// if len(attrs) > 0 {
			// 	fmt.Printf("Attributes: %+v\n", attrs)
			// 	for _, attr := range attrs {
			// 		fmt.Printf("Attribute: Key=%v, Val=%v\n", attr.Key, attr.Val)
			// 	}
			// } else {
			// 	fmt.Println("No attributes found for this token.")
			// }
			// Check if the tag is an img tag
			tagName := tokenizer.Token().Data
			if StringInSlice(tagName, tags...) {
				// Check for the presence of required attributes
				hasRelIcon := false
				hasTypeImage := false
				hasURLFavicon := false
				fmt.Println(tokenizer.Token().Attr)
				for _, attr := range tokenizer.Token().Attr {
					fmt.Printf("Tag: %v, Key: %v, Val: %v", tagName, attr.Key, attr.Val)
					if attr.Key == "rel" && strings.Contains(strings.ToLower(attr.Val), "icon") {
						hasRelIcon = true
					}
					if attr.Key == "type" && strings.Contains(strings.ToLower(attr.Val), "image") {
						hasTypeImage = true
					}
					if strings.ToLower(attr.Key) == "href" && strings.Contains(strings.ToLower(attr.Val), "favicon") {
						hasURLFavicon = true
					}
				}
				if hasRelIcon || hasTypeImage || hasURLFavicon {
					fmt.Printf("Found matching img tag: %s\n", tokenizer.Token().Data)
				}
			}
		}
	}
}
func printNode(node *html.Node, file *os.File) {
	if node.Type == html.ElementNode {
		fmt.Fprintf(file, "Tag: %v\n", node.Data)
		for _, attr := range node.Attr {
			fmt.Fprintf(file, "Attribute: %v = %v\n", attr.Key, attr.Val)
		}
	}

	if node.Type == html.TextNode || node.Type == html.DocumentNode {
		fmt.Fprintf(file, "Text: %v\n", node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		printNode(child, file)
	}
}
