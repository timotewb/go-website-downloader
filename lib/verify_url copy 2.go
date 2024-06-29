package lib

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"time"

// 	"golang.org/x/net/html"
// )

// type ResponseType struct {
// 	Code       int    `json:"code"`
// 	Message    string `json:"message"`
// 	Url        string `json:"url"`
// 	FaviconURL string `json:"favicon_url"`
// }

// func ValidateURL(urlString string) (ResponseType, error) {
// 	time.Sleep(1 * time.Second)
// 	// create response object
// 	var r ResponseType
// 	// check url is valid
// 	u, err := url.ParseRequestURI(urlString)
// 	r.Url = u.String()
// 	fmt.Println("--- Parse Done")
// 	// error check
// 	if err != nil {
// 		r.Code = 1
// 		r.Message = err.Error()
// 		return r, err
// 	}
// 	fmt.Println("--- Error Check Done")
// 	r.Code = 0
// 	r.Message = "Success"
// 	fmt.Println("--- ready to return")
// 	return r, err
// }

// func VerifyURL(r ResponseType) (ResponseType, *http.Response, error) {
// 	resp, err := http.Get(r.Url)
// 	if err != nil {
// 		r.Code = 2
// 		r.Message = err.Error()
// 		return r, resp, err
// 	}
// 	return r, resp, nil
// }

// func GetFavicon(r ResponseType, page *http.Response) (ResponseType, error) {

// 	// Write the response body (HTML content) to a text file
// 	bodyBytes, err := io.ReadAll(page.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = os.WriteFile("response.html", bodyBytes, 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to write response to file: %v", err)
// 	}

// 	doc, err := html.Parse(page.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		r.Code = 3
// 		r.Message = err.Error()
// 		return r, err
// 	}

// 	// Function to count child nodes
// 	var countNodes func(*html.Node, int) int
// 	countNodes = func(node *html.Node, count int) int {
// 		for c := node.FirstChild; c != nil; c = c.NextSibling {
// 			count++
// 			count = countNodes(c, count)
// 		}
// 		return count
// 	}

// 	// Counting the total number of nodes in the document
// 	numNodes := countNodes(doc, 0)

// 	fmt.Printf("Total number of nodes parsed: %d\n", numNodes)

// 	// base url
// 	u, err := url.Parse(r.Url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		r.Code = 3
// 		r.Message = err.Error()
// 		return r, err
// 	}
// 	u.Path = ""
// 	u.RawQuery = ""
// 	u.Fragment = ""

// 	// find favicon
// 	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	var favi func(*html.Node) bool
// 	var html_rel bool
// 	var html_type bool
// 	var html_href bool
// 	var html_href_url string

// 	favi = func(n *html.Node) bool {

// 		printNode(n, file)
// 		// l := []string{"link", "meta", "head"}
// 		fmt.Println(n.Data)
// 		// if StringInSlice(n.Data, l...) {
// 		fmt.Println("GetFavicon(): searching page content for favicon.")
// 		fmt.Fprint(file, "----------\n")
// 		for _, a := range n.Attr {
// 			fmt.Fprintf(file, "Key: %s\nNamespace: %s\nVal: %s\n", a.Key, a.Namespace, a.Val)

// 			if strings.ToLower(a.Key) == "rel" && strings.Contains(strings.ToLower(a.Val), "icon") {
// 				html_rel = true
// 			}
// 			if strings.ToLower(a.Key) == "type" && strings.Contains(strings.ToLower(a.Val), "image") {
// 				html_type = true
// 			}
// 			if strings.ToLower(a.Key) == "href" && strings.Contains(strings.ToLower(a.Val), "favicon") {
// 				html_href = true
// 				html_href_url = a.Val
// 			}
// 		}
// 		fmt.Printf("html_rel: %v, html_type: %v, html_href: %v", html_rel, html_type, html_href)
// 		if html_rel && html_type && html_href {
// 			fmt.Printf("icon: %v\n", html_href_url)
// 			r.FaviconURL = html_href_url
// 			html_rel = false
// 			html_type = false
// 			html_href = false
// 			return true
// 		}
// 		fmt.Fprint(file, "----------\n")
// 		// }
// 		// traverses the HTML of the webpage from the first child node
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			numNodes := countNodes(c, 0)
// 			fmt.Printf("   Total number of subnodes parsed: %d\n", numNodes)
// 			_ = favi(c)
// 		}
// 		return false
// 	}
// 	_ = favi(doc)
// 	if r.FaviconURL == "" {
// 		fmt.Println("GetFavicon(): trying fall back url.")
// 		resp, err := http.Get(CreateURL(u.String(), "favicon.ico"))
// 		if err != nil {
// 			fmt.Printf("Failed to perform request: %v\n", err)
// 		} else {
// 			defer resp.Body.Close()
// 			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
// 				if strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "image") {
// 					r.FaviconURL = CreateURL(u.String(), "favicon.ico")
// 				}
// 			}
// 		}
// 	}

// 	if r.FaviconURL == "" {
// 		fmt.Println("Favicon not found!")
// 	}
// 	return r, err
// }
// func printNode(node *html.Node, file *os.File) {
// 	if node.Type == html.ElementNode {
// 		fmt.Fprintf(file, "Tag: %v\n", node.Data)
// 		for _, attr := range node.Attr {
// 			fmt.Fprintf(file, "Attribute: %v = %v\n", attr.Key, attr.Val)
// 		}
// 	}

// 	if node.Type == html.TextNode || node.Type == html.DocumentNode {
// 		fmt.Fprintf(file, "Text: %v\n", node.Data)
// 	}

// 	for child := node.FirstChild; child != nil; child = child.NextSibling {
// 		printNode(child, file)
// 	}
// }
