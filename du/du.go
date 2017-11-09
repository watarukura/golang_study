package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func FileList(d string) []os.FileInfo {
	dorf, _ := ioutil.ReadDir(d)
	var fileinfos []os.FileInfo
	for i := 0; i < len(dorf); i++ {
		if dorf[i].IsDir() {
			dir := d + "/" + dorf[i].Name()
			infos := FileList(dir)
			for k := 0; k < len(infos); k++ {
				fileinfos = append(fileinfos, infos[k])
			}
		} else {
			fileinfos = append(fileinfos, dorf[i])
		}
	}
	return fileinfos
}

func main() {
	var dir string

	flag.Parse();
	dirs := flag.Args()

	if len(dirs) == 0 {
		dir = "./"
	} else {
		dir = dirs[0]
	}

	files := FileList(dir)
	count := len(files)

	var size int64
	for i := 0; i < len(files); i++ {
		size += files[i].Size()
	}
	fmt.Printf("%d files, %d bytes", count, size)
}
