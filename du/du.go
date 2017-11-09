package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func FilesCount(d string) int {
	files, _ := ioutil.ReadDir(d)
	return len(files)
}

func SizeOf(d string) int64 {
	files, _ := ioutil.ReadDir(d)
	var sum int64
	for i := 0; i < len(files); i++ {
		//fmt.Println(files[i].Name())
		if IsDir(files[i]) {
			//fmt.Println(d+files[i].Name())
			sum += SizeOf(d+files[i].Name())
			fmt.Println(sum)
		} else {
			sum += files[i].Size()
			fmt.Println(sum)
		}
	}
	return sum
}

func FilesList(d string) []os.FileInfo {
	dorf, _ := ioutil.ReadDir(d)
	var fileinfos []os.FileInfo
	j := 0
	for i := 0; i < len(dorf); i++ {
		if dorf[i].IsDir() {
			dir := d + "/" + dorf[i].Name()
			infos := FilesList(dir)
			for k := 0; k < len(infos); k++ {
				fileinfos[j] = infos[k]
				j++
			}
		} else {
			fileinfos[j] = dorf[i]
			j++
		}
	}
	return fileinfos
}

func main() {
	flag.Parse();
	dirs := flag.Args()
	fmt.Println(dirs)
	dir_files := 0
	var dir_sizes int64
	for i := 0; i < len(dirs); i++ {
		files, _ := ioutil.ReadDir(dirs[i])
		for j := 0; j < len(files); j++ {
			if files[j].IsDir() {

			}
		}
		dir_files += FilesCount(dirs[i])
		dir_sizes += SizeOf(dirs[i])
	}
	fmt.Printf("%d files, %d bytes", dir_files, dir_sizes)
}
