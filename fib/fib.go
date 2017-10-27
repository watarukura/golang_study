package main

import (
	"os"
	"fmt"
	"strconv"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
	}
	fmt.Printf("%d", fib(i))
}
