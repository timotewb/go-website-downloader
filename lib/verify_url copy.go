package lib

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"regexp"
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

// 	faviconExts := []string{".ico", ".png"}

// 	doc, err := html.Parse(page.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		r.Code = 3
// 		r.Message = err.Error()
// 		return r, err
// 	}
// 	defer page.Body.Close()

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

// 	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// find favicon.ico
// 	var favi func(*html.Node) bool
// 	favi = func(n *html.Node) bool {
// 		l := []string{"link", "meta"}
// 		rgx, _ := regexp.Compile(`favicon[^/]+.(ico|png)`)
// 		if err != nil {
// 			log.Fatalf("Failed to write to file: %v", err)
// 		}
// 		if StringInSlice(n.Data, l...) {
// 			for _, a := range n.Attr {
// 				if strings.Contains(a.Val, "favicon.ico") {
// 					fmt.Println("01")
// 					// this code is replicated - fix this
// 					if a.Val == "favicon.ico" {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + "/" + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "//") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "/") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "http") {
// 						r.FaviconURL = a.Val
// 						return false
// 					} else {
// 						fmt.Println(" --- Build new rule for GetFavicon() 01.")
// 					}
// 				} else if strings.Contains(a.Val, "favicon.png") {
// 					fmt.Println("02")
// 					if a.Val == "favicon.png" {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + "/" + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "//") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "/") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "http") {
// 						r.FaviconURL = a.Val
// 						return false
// 					} else {
// 						fmt.Println(" --- Build new rule for GetFavicon() 02.")
// 					}
// 				} else if rgx.MatchString(a.Val) {
// 					fmt.Println("03")
// 					// this code is slightly different - Yay!
// 					if strings.HasPrefix(a.Val, "//") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "/") {
// 						r.FaviconURL = strings.TrimSuffix(u.String(), "/") + a.Val
// 						return false
// 					} else if strings.HasPrefix(a.Val, "http") {
// 						r.FaviconURL = a.Val
// 						return false
// 					} else {
// 						fmt.Println(" --- Build new rule for GetFavicon() 03.")
// 					}
// 				} else if strings.Contains(strings.ToLower(a.Val), "/favicon/") {
// 					fmt.Println("04", a.Val)
// 					for _, ending := range faviconExts {
// 						if strings.HasSuffix(a.Val, ending) {
// 							if strings.HasPrefix(a.Val, "//") {
// 								r.FaviconURL = a.Val
// 								return false
// 							} else {
// 								r.FaviconURL = CreateURL(u.String(), a.Val)
// 								return false
// 							}
// 						}
// 					}
// 				} else {
// 					resp, err := http.Get(CreateURL(u.String(), "favicon.ico"))
// 					if err != nil {
// 						fmt.Printf("Failed to perform request: %v\n", err)
// 					} else {
// 						defer resp.Body.Close()
// 						if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
// 							fmt.Println("00")
// 							fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
// 							r.FaviconURL = CreateURL(u.String(), "favicon.ico")
// 							return false
// 						} else {
// 							return true
// 						}
// 					}
// 				}
// 			}
// 		}
// 		// traverses the HTML of the webpage from the first child node
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			if !favi(c) {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	_ = favi(doc)
// 	return r, err
// }

// respBody, err := os.OpenFile("resBody.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
// if err != nil {
// 	log.Fatalf("Failed to open file: %v", err)
// }
// defer respBody.Close()

// bodyBytes, err := io.ReadAll(page.Body)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Fprint(respBody, string(bodyBytes))
