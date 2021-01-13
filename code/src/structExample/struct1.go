package structExample

import "fmt"

func Run(){
	var xiaoming = struct {
		name string
		age int
		hobby []string
		education map[string]bool
	} {
		name: "xiaoming",
		age: 12,
		hobby: []string{"basketball", "football"},
		education: map[string]bool{
			"primary" : true,
			"medium" : true,
			"senior" : false,
		},
	}

	xiaoming.age++
	fmt.Println(xiaoming)
	fmt.Println(xiaoming.name)

	var agePoint = &xiaoming.age
	*agePoint = 100
	fmt.Println(xiaoming.age)	// 100
}
