package download

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/url"
	neturl "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func applyActions(node *html.Node, dm *downloadManagerType) {
	switch node.Data {
	case "a":
		aAction(node, dm)
	}
}

func writeOutPage(n *html.Node, lt linkType, dm *downloadManagerType) {
	log.Printf("writeOutPage('%v')", lt.ValOriginal)
	// create dirs
	dirName := filepath.Dir(lt.ValNew)
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	// write file
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	err = os.WriteFile(lt.ValNew, []byte(buf.String()), 0644)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	} else {
		// update linkType
		updateDMLT(lt, dm)
	}
}

func updateDMLT(lt linkType, dm *downloadManagerType) {
	for i, l := range dm.Links {
		if l.ValNew == lt.ValNew {
			dm.Links[i].WrittenOut = true
			return
		}
	}
}

func aAction(node *html.Node, dm *downloadManagerType) {
	// find the href attribute and get tidy link
	for i, attr := range node.Attr {
		if attr.Key == "href" {
			// check if we have already actioned the path
			var lt linkType
			lt = checkLinkTypes(node, attr, dm)
			if lt.IsEmpty() {
				// we have not actioned this path
				lt = addLinkType(node, attr, dm, "page")
			}
			if !lt.IsEmpty() {
				// Create a new attribute with the modified value and update
				// overwrite original attribute
				newNodeAttr := html.Attribute{Key: lt.Attr, Namespace: attr.Namespace, Val: lt.ValNew}
				node.Attr[i] = newNodeAttr

				// add original val
				newNodeAttr = html.Attribute{Key: "original_" + lt.Attr, Namespace: attr.Namespace, Val: lt.ValOriginal}
				node.Attr = append(node.Attr, newNodeAttr)

				// add get url val
				newNodeAttr = html.Attribute{Key: "geturl_" + lt.Attr, Namespace: attr.Namespace, Val: lt.GetURL}
				node.Attr = append(node.Attr, newNodeAttr)
			}
		}
	}
}

func addLinkType(node *html.Node, attr html.Attribute, dm *downloadManagerType, kind string) linkType {

	var lt linkType
	lt.WrittenOut = false
	if kind == "page" {
		lt.ValNew = filepath.Join(dm.RootDir, kind, generateRandomString(10, dm)+".html")
	}

	url := attr.Val

	if strings.Contains(url, "#") {
		return lt
	}

	// url starting with rootURL, i.e. internal link
	if strings.HasPrefix(url, dm.RootURL) {

		lt.Data = node.Data
		lt.Attr = attr.Key
		lt.ValOriginal = url
		lt.GetURL = url

		dm.Links = append(dm.Links, lt)
		return lt
	}

	// url starting with / and not root (e.g. /about)
	realitiveLink, _ := regexp.MatchString(`^\/[a-zA-Z0-9]+`, url)
	if realitiveLink {
		s, err := neturl.JoinPath(dm.RootURL, url)
		if err != nil {
			log.Print(err)
		}
		lt.Data = node.Data
		lt.Attr = attr.Key
		lt.ValOriginal = url
		lt.GetURL = s

		dm.Links = append(dm.Links, lt)
		return lt
	}

	// url starting with // e.g. //about dfs=Double Forward Slash
	dfsLink, _ := regexp.MatchString(`^\/[a-zA-Z0-9]+`, url)
	if dfsLink {
		return lt
	}

	// check if the url contins the domain
	if strings.Contains(url, dm.Domain) {
		return lt
	}

	// url starting with http:// or https:// (used later)
	httpLink, _ := regexp.MatchString(`^http:\/\/|https:\/\/[a-zA-Z0-9]+`, url)
	if httpLink {
		return lt
	}

	if url == "/" {
		return lt
	}

	return lt
}

func checkLinkTypes(node *html.Node, attr html.Attribute, dm *downloadManagerType) linkType {
	var lt linkType
	for i, _ := range dm.Links {
		if dm.Links[i].Data == node.Data && dm.Links[i].Attr == attr.Key && dm.Links[i].ValOriginal == attr.Val {
			return dm.Links[i]
		}
		// if dm.Links[i].ValOriginal == attr.Val {
		// 	return dm.Links[i]
		// }
	}
	return lt
}

func checkVarNewUsed(varNew string, dm *downloadManagerType) bool {
	for _, l := range dm.Links {
		if varNew == l.ValNew {
			return true
		}
	}
	return false
}

func getRootDomain(inputURL string) (string, error) {
	// Parse the input URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	// Get the host part of the URL
	host := parsedURL.Host

	// Remove "www." from the host if present
	if strings.HasPrefix(host, "www.") {
		host = host[4:]
	}

	return host, nil
}

func getRootURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	rootURL := *parsedURL
	rootURL.Host = parsedURL.Hostname()
	rootURL.Scheme = parsedURL.Scheme
	rootURL.Path = "/"
	rootURL.RawQuery = ""
	rootURL.Fragment = ""

	return rootURL.String(), nil
}

// GenerateAlphanumericString generates a random string of letters (both uppercase and lowercase) and numbers of the specified length.
func generateRandomString(length int, dm *downloadManagerType) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)
	rand.Read(b)

	var sb []byte
	for _, v := range b {
		sb = append(sb, charset[v%byte(len(charset))])
	}

	checkName := true
	for checkName {
		if !checkVarNewUsed(string(sb), dm) {
			checkName = false
		}
	}

	return string(sb)
}
