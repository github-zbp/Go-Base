package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

// 获取一个url的内容并且打印到标准输出
func FetchPrint(url string) {
	content, err := Fetch(url)

	// 如果有错误，则将错误输出到标准错误（其实还是输出到屏幕上，但是相比于标准输出的区别是，标准错误没有缓冲区，所以会比标准错误经常会比标准输出要更早打印出来）
	// Fprintf作用是将第三参的数据按第二参格式化之后输入到第一参指定的流中，可以是文件流（那就是写入到文件），标准输出流，标准错误流等
	if err !=nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	fmt.Fprintf(os.Stdout, "%s", content)		// 这里用 fmt.Printf("%s", content) 也可以，因为Printf默认就是输出到标准输出流，也就是打印到屏幕上
}

// 获取一个url的内容并返回响应内容（字符串）
func Fetch(url string) (string, error){
	var content string

	body, err := FetchBody(url)

	if err != nil {
		return content, err
	}

	// ioutil.ReadAll通过流的方式读取响应中的字节流内容返回接收到的所有字节流（[]byte字节切片类型），需要传入一个io.Reader类型的变量，而resp.Body是io.ReadCloser类型（继承了Reader类型）
	b, err := ioutil.ReadAll(body)		// resp的Body成员包括一个可读的服务器响应流
	body.Close()		// 不论有没有发生错误，都得关闭resp的Body流，防止资源泄露

	if err != nil {
		return content, fmt.Errorf("链接 %s :io读取响应数据失败: %v", url, err)
	}

	content = string(b)		// 将字节流转为字符串
	return content, nil
}

// 该函数返回resp.Body
func FetchBody(url string) (body io.ReadCloser, err error) {
	// 获取一个url的响应, 返回一个Response类型变量
	resp, err := http.Get(url)

	if err != nil {
		// fmt.Errorf 返回一个error错误类型的变量，由于这里要求返回一个error类型所以不能使用Sprintf
		return body, fmt.Errorf("获取链接失败: %v", err)
	}

	//if string(string(resp.StatusCode)[0]) != "2" {
	// strconv.Itoa(x int)的作用是将一个int类型的整型转化为字符串型的数字。而 string() 如果传入一个整型会转为一个对应的ASCII
	// 所以 正确的数字转字符串的方式是 strconv.Itoa() 而不是 stirng()
	// 而且假如a是字符串类型, a[0]不是一个字符串，而是一个uint8， 所以 判断a[0]是否等于字符a不应该用 a[0] == "a" ，而是用 a[0] == 'a' 在Go中双引号是字符串类型，而单引号是rune类型
	if strconv.Itoa(resp.StatusCode)[0] != '2' {
		return body, fmt.Errorf("获取链接返回状态不正常: %s | %s | %s", resp.Status, string(string(resp.StatusCode)[0]), string(resp.StatusCode)[0])
	}

	// 如果请求成功，则返回响应流对象（关闭响应流要在读取内容之后执行）
	return resp.Body, err
}

// 请求一个url并将内容写入到文件中,返回文件的绝对路径
func FetchToFile(url string) (filePath string, err error){
	resp, err := http.Get(url)
	if err != nil {
		return filePath, err
	}

	defer resp.Body.Close()

	// 获取url的文件名
	fileName := path.Base(resp.Request.URL.Path)
	if fileName == "/"{
		fileName = "index.html"
	}

	// 将数据读到文件
	relatePath := "./" + fileName
	f, err := os.Create(relatePath)
	if err != nil {
		return filePath, err
	}


	_, err = io.Copy(f, resp.Body)		// 返回写入文件的字节数

	defer closeFile(f, &err)

	// 如果没有发生错误
	filePath,_ = filepath.Abs(relatePath)

	return filePath, err
}

func closeFile(f *os.File, err *error) {
	// 关闭文件
	if errClose := f.Close(); *err == nil {
		*err = errClose
	}
}