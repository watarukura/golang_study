package main

import (
	"encoding/json"
	"net/http"
	"log"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Get(url string) (*Page, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	page := &Page{}
	var f func(node *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			page.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			if isDescription(n.Attr) {
				for _, attr := range n.Attr {
					if attr.Key == "content" {
						page.Description = attr.Val
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return page, nil
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		rawurls := r.Form["url[]"]
		for i := 0; i < len(rawurls); i++ {
			if rawurls[i] == "" {
				http.Error(w, "url not specified", http.StatusBadRequest)
				return
			}
			p, err := Get(rawurls[i])
			if err != nil {
				http.Error(w, "request failed", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			enc := json.NewEncoder(w)
			if err := enc.Encode(p); err != nil {
				http.Error(w, "encoding failed", http.StatusInternalServerError)
				return
			}

		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
