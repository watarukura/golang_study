package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func FileList(d string) []os.FileInfo {
	var wg sync.WaitGroup

	dorf, _ := ioutil.ReadDir(d)
	var fileinfos []os.FileInfo
	for i := 0; i < len(dorf); i++ {
		wg.Add(1)
		if dorf[i].IsDir() {
			dir := d + "/" + dorf[i].Name()
			infos := FileList(dir)
			for _, info := range infos {
				fileinfos = append(fileinfos, info)
			}
		} else {
			fileinfos = append(fileinfos, dorf[i])
		}
		wg.Done()
	}
	wg.Wait()
	return fileinfos
}

func main() {
	var wg sync.WaitGroup

	flag.Parse();
	dirs := flag.Args()

	// 指定がない場合はカレントディレクトリ配下に対して実行
	if len(dirs) == 0 {
		dirs = append(dirs, "./")
	}

	var size int64 // os.Fileinfo.Size() の返り値はint64なのでそのまま使う
	var count int

	for _, dir := range dirs {
		wg.Add(1)
		files := FileList(dir)
		count += len(files)
		for _, file := range files {
			size += file.Size()
		}
		wg.Done()
	}
	wg.Wait()
	fmt.Printf("%d files, %d bytes", count, size)
}
