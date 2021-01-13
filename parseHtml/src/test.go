package main

import (
	//"encoding/json"
	"fetch"

	//"encoding/json"
	"fmt"
	//"io/ioutil"

	//"io/ioutil"
	//"parseImprove"
)

func main() {
	// 获取一个域名下所有链接,可以有多个入口文件
	//startUrls := []string {"http://www.zbpblog.com"}
	//graph := map[string][]string {}
	//res := parseImprove.GetDomainDepthLinks(startUrls,graph)
	//
	//fmt.Printf("%#v\n\n\n\n\n\n\n\n", graph)
	//
	////将graph存到文件，这样下次想深度遍历这个域名的时候就不用再构建一次图了
	//graphJson,_ := json.Marshal(graph)
	//ioutil.WriteFile("1.json",graphJson, 0777)
	//
	//fmt.Printf("%#v", res)
	//fmt.Printf("%d", len(res))

	//workList := []string {"http://www.zbpblog.com"}
	//parseImprove.GetDomainWidthLinks(parseImprove.GetLinks, workList)
	//fmt.Println(workList)

	fmt.Println(fetch.FetchToFile("http://www.zbpblog.com/blog-94.html"))
}