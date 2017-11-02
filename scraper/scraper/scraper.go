package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	enc := json.NewEncoder(w)
	if err := enc.Encode(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
