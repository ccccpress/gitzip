package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// 获取当前时间作为文件名
	now := time.Now().Format("2006_01_02_15_04_05")

	const dir = `.\`
	f, err := ioutil.ReadDir(dir)
	checkerr(err)
	// 如果文件夹已存在会自动忽视
	fzip, err := os.Create("gitzip" + now)
	checkerr(err)
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		// 不会遍历文件夹！！！
		if file.IsDir() || (len(file.Name()) > 5 && file.Name()[:6] == "gitzip") {
			continue
		}
		fw, err := w.Create(file.Name())
		checkerr(err)
		filecontent, err := ioutil.ReadFile(dir + file.Name())
		checkerr(err)
		n, err := fw.Write(filecontent)
		checkerr(err)
		fmt.Println(now, n)
	}
}
func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
