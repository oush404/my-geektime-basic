package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name  string `json:"name"`
	Age   int
	sex   string
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

type StuRead struct {
	Name interface{} `json:"name"`
	Age  interface{}
	//High  interface{}
	sex interface{}
	//Class interface{} `json:"class"`
	Class *Class `json:"class"`
	//Test  interface{}
}

func main() {
	stu := Stu{
		Name: "张三",
		Age:  18,
		sex:  "男",
	}
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(jsonStu)
	fmt.Println(string(jsonStu))

	data := "{\"name\":\"张三22\",\"Age\":18,\"class\":{\"Name\":\"1班\",\"Grade\":3}}"
	str := []byte(data)
	stu2 := new(StuRead)
	err = json.Unmarshal(str, &stu2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(stu2)
	fmt.Println(stu2.Age)
	fmt.Println(stu2.Class)
	fmt.Println(stu2.Class.Grade)
	fmt.Println(stu2.Class.Name)
}
