package main

import "fmt"

func main() {
	//map1 := make(map[string]string)		// 初始化一个key和value都是字符串的哈希表
	//map1["name"] = "zbp"	// 通过下表的方式添加元素，这里不能用单引号，单引号表示字符类型，双引号才是字符串类型！
	//map1["age"] = "24"		// 哈希表的value必须是相同类型，所以这里不能是数字24，只能是字符串24
	//
	//fmt.Println(map1)

	// example2
	//map2 := map[string]string {
	//	"name" : "zbp",
	//	"age" : "24",		// 这里必须要又逗号，反则会报错
	//}
	//fmt.Println(map2)

	// example3
	//map3 := map[string]string{}
	//map3["name"] = "zbp"
	//map3["age"] = "24"
	//fmt.Println(map3)
	//fmt.Println(map3["hobby"])		// 获取哈希表中不存在的元素不会报错，会默认返回value类型的零值

	// example4
	//map4_score := map[string]int{		// 班级数学平均分
	//	"class1" : 86,
	//	"class2" : 90,
	//}
	//fmt.Println(map4_score)
	//delete(map4_score, "class2")
	//fmt.Println(map4_score)
	//delete(map4_score, "class4")	// 删除不存在的key不会报错
	//fmt.Println(map4_score)

	// example5
	//map5 := map[string]string {
	//	"name" : "zbp",
	//	"age" :"24",
	//	"hobby" : "programme",
	//}
	//
	//for key,value := range map5 {
	//	fmt.Printf("key : %s, value : %s\n", key, value)
	//}

	// example6
	//var map6 map[string]int
	//map6 = map[string]int{}
	//map6["num1"] = 1
	//fmt.Println(map6)

	// example7
	//map7 := make(map[string]string)
	//map7["name"] = "zbp"
	//age, ok := map7["age"]
	//fmt.Println(age, ok)
	//
	//if _, ok := map7["age"]; !ok{
	//	fmt.Println("Key of age does not exist!")
	//}

	// example8
	//map8 := map[string]string{
	//	"k1" : "v1",
	//	"k2" : "v2",
	//}
	//map9 := map[string]string{
	//	"k1" : "v1",
	//	"k2" : "v2",
	//}
	//fmt.Println(equal(map8, map9))

	// example9
	//slice1 := []string {"a", "b", "c"}
	//slice2 := []string {"e", "f"}
	//map10 := map[string]bool{}
	//map10[k(slice1)] = true
	//map10[k(slice2)] = false
	//fmt.Println(map10)

	map11 := map[string]map[string]int{}	// 创建一个vlaue是map的变量，这个变量记录不同班级每个学生的数学成绩
	class1 := map[string]int{
		"小明": 100,
		"小红": 98,
	}
	class2 := map[string]int{
		"小黑" : 99,
		"小白" : 60,
	}
	map11["class1"] = class1
	map11["class2"] = class2
	fmt.Println(map11)

}

func equal(x, y map[string]string) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if _, ok := y[k]; !ok || y[k] != xv {
			return false
		}
	}

	return true
}

func k(slice []string) string {
	return fmt.Sprintf("%q", slice)
}