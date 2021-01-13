package funcExample
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"golang.org/x/net/html"
//)
//
//func Run() {
//	// 请求一个url
//	url := "http://zbpblog.com"
//	htmlText := fetch(url)
//
//	// 开始解析html
//	nodes, err := html.Parse(byte(htmlText))
//}
//
//// 爬取一个url，返回响应内容的string格式
//func fetch(url string) string {
//	var (
//		resp *http.Response
//		err error
//	)
//	if resp, err = http.Get(url); err != nil{
//		// Fprintf() 根据 format格式说明符将内容格式化并以流的形式写入文件（或者其他流），返回内容是写入的字节与错误。
//		// 这里是将格式化后的信息写入到标准错误中，对标准输出、标准输入和标准错误感兴趣的朋友可以参考 https://blog.csdn.net/qingzhuyuxian/article/details/80391143
//		fmt.Fprintf(os.Stderr, "fetch:%v\n", err )
//		os.Exit(1)	// 退出程序
//	}
//
//	// 读取响应内容，读取得到的是字节流格式的数据
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "fetch reading:%v\n", err)
//	}
//
//	respText := fmt.Sprint("%s", b)
//
//	return respText
//}