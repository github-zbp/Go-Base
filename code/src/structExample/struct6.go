package structExample

import "fmt"

func Run6() {
	type person struct {
		name string
		age int
	}

	type student struct {
		person
		school string
		class string
	}

	type worker struct {
		person
		job string
		level int
	}

	//zbp := new(worker)
	//zbp.name = "zbp"
	//zbp.age = 24
	//zbp.job = "programmer"
	//zbp.level = 3
	//fmt.Println(*zbp)
	//fmt.Println(zbp.person)

	//zbp := worker {
	//	person: person {
	//		name: "zbp",
	//		age: 24,
	//	},
	//	job: "programmer",
	//	level: 3,
	//}

	zbp := worker {
		person {"zbp", 24,}, "programmer", 3,
	}

	//fmt.Println(zbp)
	fmt.Printf("%#v", zbp)
}
