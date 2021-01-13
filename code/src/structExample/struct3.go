package structExample

import "fmt"

//type employee1 struct {
//	name string
//	level int
//}
//
//type employee2 struct {
//	name string
//	level int
//}

func Run3() {
	var zbp *Person = &Person{
		name:"zbp",
		age:24,
	}

	zbp.age = 25	// 通过结构体指针操作结构体
	fmt.Println(*zbp)
	(*zbp).age = 26	// 直接操作结构体本身
	fmt.Println(*zbp)

}

func TestStructType() {
	e1 := struct {
		name string
		level int
	} {}
	e2 := struct {
		name string
		level int
	} {}
	fmt.Println(e1==e2)
}
