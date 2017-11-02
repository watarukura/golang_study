package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"flag"
	"strings"
)

func main() {
	var (
		method  string
		url     string
		head    string
		payload string
	)

	flag.StringVar(&method, "X", "GET", "http method")
	flag.StringVar(&head, "head", "", "http header")
	flag.StringVar(&head, "H", "", "http header")
	flag.StringVar(&payload, "d", "", "payload url-encode")
	flag.StringVar(&payload, "data-urlencode", "", "payload url-encode")
	flag.Parse();
	url = flag.Args()[0]
	heads := strings.Fields(head)
	head_key := strings.Replace(heads[0], ":", "", -1)
	head_val := heads[1]

	req, err := http.NewRequest(method, url, nil)
	if head != "" {
		req.Header.Set(head_key, head_val)
	}
	if payload != "" {
		req.PostFormValue(payload)
	}
	if err != nil {
		fmt.Println(req.Body)
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(body))

}
