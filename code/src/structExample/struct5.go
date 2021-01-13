package structExample

import "fmt"

func Run5() {
	zbp := employee{"zbp", 3}
	zbp2 := employee{"zbp", 3}
	fmt.Println(zbp == zbp2)
	fmt.Println(zbp.name == zbp2.name && zbp.level == zbp2.level)	// 等价于fmt.Println(zbp == zbp2)
}
