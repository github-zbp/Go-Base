package funcExample

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetUrl(url string) (content []byte, err error){
	resp, err := http.Get(url)
	defer connectClose(resp)
	if err != nil {
		return content, fmt.Errorf("%s get error: %v, not text/html", url, err)
	}

	// 获取网页的content-type，如果不为 text/html 则直接返回；是text/html的格式才获取响应内容
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return content, fmt.Errorf("%s has type %s, not text/html",url, ct)
	}

	// 将响应流数据写入到content中
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return content, fmt.Errorf("%s read error: %v", url, err)
	}

	return content, err
}


func GetUrlWithoutDefer(url string) (content []byte, err error){
	resp, err := http.Get(url)
	if err != nil {
		return content, fmt.Errorf("%s get error: %v, not text/html",url, err)
	}

	// 获取网页的content-type，如果不为 text/html 则直接返回；是text/html的格式才获取响应内容
	ct := resp.Header.Get("Content-Type")
	resp.Body.Close()
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return content, fmt.Errorf("%s has type %s, not text/html",url, ct)
	}

	// 将响应流数据写入到content中
	content, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return content, fmt.Errorf("%s read error: %v", url, err)
	}

	return content, err
}

func connectClose(resp *http.Response) {
	fmt.Println("关闭连接")
	resp.Body.Close()
}