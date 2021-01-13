package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetUrl(url string) (cont interface{}, err error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	cont, err = ioutil.ReadAll(resp.Body)
	return cont, err
}

// 写入文件
func Save(url, dir string, cont []byte) {
	if strings.Contains(url, "php"){
		return
	}

	fn := strings.TrimRight(strings.TrimRight(url, "/"), ".html")
	fn = filepath.Base(fn) + ".html"
	dir = strings.TrimRight(dir, "/") + "/"
	fp := dir + fn

	_,err := os.Stat(dir)
	if err != nil && os.IsNotExist(err){
		os.Mkdir(dir, 0777)
	}

	f, err := os.OpenFile(fp, os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("写入文件 %s 失败, 原因: %s \n", fp, err)
	}
	defer f.Close()
	f.Write([]byte(cont))
}