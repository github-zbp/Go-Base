package jsonExample

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name	string
	Year	int `json:"released"`
	Color	bool `json:color, omitempty`	//是否为彩色电影
	Actors []string
	level 	int
}

func Run() {
	var movies []Movie // 定义一个电影列表，里面的类型是Movie结构体类型，相当于python中一个列表里面放多个字典
	movies = []Movie{
		Movie{
			Name:   "阿甘正传",
			Year:   1994,
			Color:  true,
			Actors: []string{"A", "B", "C"},
			level: 9,
		},
		 Movie{
			Name:   "霸王别姬",
			Year:   1993,
			Color:  true,
			Actors: []string{"D", "E"},
		},
		Movie{
			Name:   "乱世佳人",
			Year:   1939,
			Color:  false,
			Actors: []string{"F", "G"},
		},
		Movie{
			Name:   "罗马假日",
			Year:   1953,
			Actors: []string{"H", "I"},
		},
}

	res, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	resFormat, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", res)
	fmt.Printf("%s\n", resFormat)

	lifetimes := Movie {
		Name:   "活着",
		Year:   1994,
		Actors: []string{"Z", "X"},
	}

	res, err = json.Marshal(lifetimes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", res)
}
