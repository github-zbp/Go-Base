package spider

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/* 爬虫 包含请求器 下载器 解析器 */

// 请求器
func Crawl(url string) (cont []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect: '%s' because of %s", url, err)
	}

	defer resp.Body.Close()
	cont, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Failed to get url content: '%s' because of %s", url, err)
	}
	return cont, err
}

// 下载器
func Download(cont []byte, url, htmlDir string) error {
	if filepath.Ext(url) != ".html"{

		return nil
	}
fmt.Println("预备下载")
	if isDirNotExist(htmlDir) {
		os.Mkdir(htmlDir, 0777)
	}
//fn := md5_encrypt(url) + ".html"
	fn := filepath.Base(url)
	fp := strings.TrimRight(htmlDir, "/") + "/" + fn
fmt.Println("准备下载" + fp)
	return writeFile(cont, fp)
}

// 判断目录是否存在
func isDirNotExist(htmlDir string) bool {
	_, err := os.Stat(htmlDir)
	if err != nil && os.IsNotExist(err){
		return true
	}
	return false
}

// 对url加密，用于生成一个html文件名
func md5_encrypt(str string) (encrypt_str string){
	encrypt := md5.New()
	io.WriteString(encrypt, str)
	return fmt.Sprintf("%x", encrypt.Sum(nil))
}

// 将内容写入到一个文件
func writeFile(cont []byte, fp string) error {
	f, err := os.OpenFile(fp, os.O_CREATE, 0777)		// 以只写方式打开文件，如果文件不存在则创建
	if err != nil{
		return err
	}

	defer f.Close()
	_, err = f.Write(cont)
fmt.Printf("写入内容到 %s\n", fp)
	return err
}

// 解析器（解析url）
func ParseUrl(cont []byte, domain string) (map[string]struct{}){
	urls := map[string]struct{}{}
	reg := regexp.MustCompile("(?:<a.*?href=['\"](.*?)['\"].*?>)")
	res := reg.FindAllSubmatch(cont, -1)

	for _, r := range res {
		href := string(r[1])
		if href == "javascript:;"{
			continue
		}
//fmt.Printf("href: %s", href)
		domain = strings.TrimRight(domain, "/")
		domainWithoutScheme := strings.Replace(domain, "http://", "", -1)
		domainWithoutScheme = strings.Replace(domain, "https://", "", -1)

		// 判断是否为domain域名下的url,不是则跳过
		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href,"//"){
			if !strings.Contains(href, domainWithoutScheme){
				continue
			}
		}

		// 如果获取的url是一个相对路径就要拼接为完整路径
		if !strings.Contains(href, domainWithoutScheme) {
			href = domain + "/" + strings.TrimLeft(href, "/")
		}
		urls[href] = struct{}{}		// 选用map而不是slice来存储一个页面所有的url时为了去掉这 一 个页面的重复url，这点很重要而且并不是多余的操作，它和正式爬取的去重(seen)缺一不可
	}
//fmt.Println(urls)
	return urls
}