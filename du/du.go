package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func FilesCount(d string) int {
	files, _ := ioutil.ReadDir(d)
	return len(files)
}

func SizeOf(d string) int64 {
	files, _ := ioutil.ReadDir(d)
	var sum int64
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i].Name())
		if IsFile(files[i]) {
			sum += files[i].Size()
			fmt.Println(sum)
		} else {
			SizeOf(d+files[i].Name())
		}
	}
	return sum
}

func IsFile(fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir()
}

func main() {
	flag.Parse();
	dirs := flag.Args()
	fmt.Println(dirs)
	dir_files := 0
	var dir_sizes int64
	for i := 0; i < len(dirs); i++ {
		dir_files += FilesCount(dirs[i])
		dir_sizes += SizeOf(dirs[i])
	}
	fmt.Printf("%d files, %d bytes", dir_files, dir_sizes)
}
