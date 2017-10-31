package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://example.com/")
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
