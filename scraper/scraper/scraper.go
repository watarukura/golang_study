package main

import (
	"encoding/json"
	"net/http"
	"log"

	"golang.org/x/net/html"
	"go/ast"
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
	var f = func(node *html.Node)

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawurl := r.FormValue("url")
		if rawurl == "" {
			http.Error(w, "url not specified", http.StatusBadRequest)
			return
		}
		p, err := Get(rawurl)
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
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
