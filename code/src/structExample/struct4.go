package structExample

import "fmt"

type employee struct {
	name string
	level int
}

func Run4() {
	// 方式1：
	zbp := employee{"zbp", 3}

	// 方式2：
	weiwei := employee{name:"weiwei", level:10}

	fmt.Println(zbp, weiwei)
}
