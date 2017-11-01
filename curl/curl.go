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
		method string
		url    string
		head   string
	)

	flag.StringVar(&method, "X", "GET", "http method")
	flag.StringVar(&head, "head", "", "http header")
	flag.StringVar(&head, "H", "", "http header")
	flag.Parse();
	url = flag.Args()[0]
	heads := strings.Fields(head)
	head_key := heads[0]
	head_val := heads[1]

	req, err := http.NewRequest(method, url, nil)
	if head != "" {
		req.Header.Set(head_key, head_val)
	}
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}
	fmt.Println(string(body))

}
