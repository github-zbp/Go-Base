package funcExample

import "fmt"

func Run4() {
	// 声明一个函数，在go中函数和哈希表，切片一样也是一种数据类型
	var add func(x, y int) int		// add被声明了，但是没有被赋值，因此add是一个零值函数

	res := add(1, 2)

	fmt.Println(res)
}
