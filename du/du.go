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
			for _, info := range infos {
				fileinfos = append(fileinfos, info)
			}
		} else {
			fileinfos = append(fileinfos, dorf[i])
		}
	}
	return fileinfos
}

func main() {
	flag.Parse();
	dirs := flag.Args()

	// 指定がない場合はカレントディレクトリ配下に対して実行
	if len(dirs) == 0 {
		dirs = append(dirs, "./")
	}

	var size int64 // os.Fileinfo.Size() の返り値はint64なのでそのまま使う
	var count int

	for _, dir := range dirs {
		files := FileList(dir)
		count += len(files)
		for _, file := range files {
			size += file.Size()
		}
	}
	fmt.Printf("%d files, %d bytes", count, size)
}
