package jsonExample

import (
	"encoding/json"
	"fmt"
)

type Title struct {
	Name string
	Year int `json:"released"`
}

func Run2() {
	var movies_json = []byte(`[{"Name":"阿甘正传","released":1994,"Color":true,"Actors":["A","B","C"]},{"Name":"霸王别姬","released":1993,"Color":true,"Actors":["D","E"]},{"Name":"乱世佳人","released":1939,"Color":false,"Actors":["F","G"]},{"Name":"罗马假日","released":1953,"Color":false,"Actors":["H","I"]}]`)
	var movies []Movie		// 接收所有电影的所有成员字段
	var titles []Title		// 接收所有电影的仅Name和Year成员字段

	if err := json.Unmarshal(movies_json, &movies); err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(movies_json, &titles); err != nil {
		fmt.Println(err)
	}

	fmt.Println(movies)
	fmt.Println(titles)
}
