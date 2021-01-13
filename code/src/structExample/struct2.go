package structExample

import "fmt"

type Person struct {
	name string
	age int
	hobby []string
	education map[string]bool
}

func Run2() {
	var xiaoming = Person {
		name: "xiaoming",
		age: 12,
		hobby: []string{"basketball", "football"},
		education: map[string]bool{
			"primary": true,
			"medium": true,
			"senior": false,
		},
	}

	var xiaohong = Person {
		name: "xiaohong",
		age: 18,
		hobby: []string{"drawing", "dancing"},
		education: map[string]bool{
			"primary": true,
			"medium": true,
			"senior": true,
		},
	}

	fmt.Println(xiaoming, xiaohong)
}
