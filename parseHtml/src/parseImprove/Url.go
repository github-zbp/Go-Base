package parseImprove

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	u "net/url"
	"os"
	"strings"
)

// 获取一个页面下所有的url
func GetLinks(url string) (urls []string) {
	// 请求一个url
	//body, err := fetch.FetchBody(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 解析html
	topNode, _ := html.Parse(resp.Body)
	resp.Body.Close()

	// collectLinks 方法只是遍历节点的前置操作
	collectLinks := func (n *html.Node) {
		if n.Data == "a" {
			for _, attr := range n.Attr {
				// 如果这个url不是这个域名下的url，就不添加到urls中
				if attr.Key == "href" && !strings.HasPrefix(attr.Val, "javascript"){
					var urlInfo *u.URL
					var err_parse_url error
					fullUrl, _ := resp.Request.URL.Parse(attr.Val)	// 获取完整的url
					if urlInfo, err_parse_url = u.Parse(fullUrl.String()); err_parse_url == nil && strings.Contains(urlInfo.Host, "zbpblog.com"){
						urls = append(urls, fullUrl.String())	// urls是环境变量，会保存状态
					}
					break
				}
			}
		}
	}

	// 遍历所有节点，并且记录下节点中的链接到urls中
	ForeachNode(topNode, collectLinks, nil)
	return urls
}

// 深度优先算法获取一个域名下所有的urls
// startUrls是入口url，可以有多个; 第二参是一个图，要求调用的时候传入一个空图nil
func GetDomainDepthLinks(startUrls []string, graph map[string][]string) []string{
	if graph == nil {
		graph = map[string][]string {}
	}

	seenDepth := map[string]bool{}

	// 首先我们需要构建一个图
	var createGraph func (urls []string, graph map[string][]string)
	createGraph = func (urls []string, graph map[string][]string) {
		for _, url := range urls {
			//url = strings.Trim(url, "/")	  // 去除首尾两端的/
			if _, ok := seenDepth[url]; !ok {
				seenDepth[url] = true		// 记录访问过的url
				innerUrls := getUniqueUrls(GetLinks(url), seenDepth)	// GetLinks的返回值去重
				fmt.Printf("%s 下所有的url有：%#v \n\n", url, innerUrls)
				graph[url] = innerUrls
				createGraph(innerUrls, graph)
			}
		}
	}

	createGraph(startUrls, graph)

	// 构建图之后，再用深度优先算法遍历（或者你可以注释掉下面这行，不在GetDomainDepthLinks中遍历，而是在GetDomainDepthLinks外遍历，GetDomainDepthLinks只负责构建图也可以）
	return IterGraphByDepth(graph)
}

func getUniqueUrls(urls []string, seen map[string]bool) (unique_urls []string){
	for _, url := range urls {
		if _, ok := seen[url]; !ok {
			unique_urls = append(unique_urls, url)
		}
	}
	return unique_urls
}

// 遍历graph这个图中所有的表层节点
func IterGraphByDepth(graph map[string][]string) (result []string) {
	seen := map[string]bool{}

	// 对每一个表层节点进行深度遍历
	for startUrl, _ := range graph {
		result = parseGraphNode(startUrl, graph, result, seen)
	}

	return result
}

// 遍历graph这个图中的某一个节点以及该节点的所有相邻节点
// res用于记录一个按深度优先排序的url切片，res中排最前面的url就是整个域名中深度最深的url
// seen用于记录访问过的url，避免重复添加url到res中，因此res中没有重复的url
func parseGraphNode (url string, graph map[string][]string, res []string, seen map[string]bool) []string{
	if _, ok := seen[url]; !ok {
		seen[url] = true

		// 如果这个节点有相邻节点，就先深入查找这些相邻节点
		for _, innerUrl := range graph[url] {
			res = parseGraphNode(innerUrl, graph, res, seen)
		}

		// 把所有的相邻节点添加到了res后，才能将自己添加到res中
		res = append(res, url)
	}

	return res
}

// 广度优先算法获取一个域名下所有的urls
// workList既是入口urls也是我们想要的存着所有广度优先遍历的urls
func GetDomainWidthLinks(f func(startUrl string) []string, workList []string) (res []string){
	seen := map[string]bool{}
	loged := map[string]bool{}

	for len(workList) > 0 {
		for _, url := range  workList {
			if _, ok := seen[url]; !ok {
				//var innerUrls []string
				// 将所有innerUrls标记为已记录
				innerUrls := f(url)
				for _, innerUrl := range innerUrls {
					// 如果这个url没有被记录才要记录
					if _, ok := loged[innerUrl]; !ok {
						fmt.Printf("%s 下的所有未爬取过的url： %s\n", url, innerUrl)
						loged[innerUrl] = true
						//innerUrls =
						workList = append(workList, innerUrl)
					}
				}

				res = append(res, url)
			}
		}
	}

	return res
}

func compareSlice(s1 []string, s2 []string) bool {
	if len(s1) != len(s2){
		return false
	}

	for key := range s1 {
		if s1[key] != s2[key]{
			return false
		}
	}

	return true
}