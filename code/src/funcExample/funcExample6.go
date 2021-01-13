package funcExample

import "sort"

var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 对图进行深度优先遍历，返回一个存储着图中所有节点的切片
func TopSort(graph map[string][]string) (result []string){
	// 要用一个map记录已经遍历过的节点
	seen := make(map[string]bool)

	// 定义一个闭包，用于对图中每一个节点遍深度历
	var iter func(node string)
	iter = func(node string) {
		// 如果该节点遍历过，则跳过
		if ok := seen[node]; ok {
			return
		}

		seen[node] = true
		neighbors, ok := graph[node]

		// 如果一个节点没有相邻节点（体现为这个节点在graph中不存在）则把该节点放入到结果集中
		if !ok || len(neighbors) == 0 {
			result = append(result, node)
			return
		}

		// 递归遍历node的相邻节点
		for _, neighbor := range neighbors {
			iter(neighbor)
		}

		// 递归完node的所有相邻接点后才将node放入到结果集
		result = append(result, node)
		//return
	}

	// 在遍历这个图之前，我希望每次调用topSort遍历map时的访问元素顺序是相同的（由于map的无序性，map元素的遍历是无需的），所以用sort.String给graph的key排一下序
	var keys []string
	for key := range graph {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// 正式开始遍历这个图
	for _, node := range keys {
		//seen[node] = true
		iter(node)
	}

	return result
}
