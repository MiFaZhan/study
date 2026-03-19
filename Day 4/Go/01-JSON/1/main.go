package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Gender string
}

func main() {
	p := Person{"米法展", "男"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("json err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	b, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Printf("json err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	//struct tag
	student := make(map[string]interface{})

	student["name"] = "米法展"
	student["gender"] = "男"
	student["age"] = 18
	c, err := json.Marshal(student)
	if err != nil {
		fmt.Printf("json err:%v\n", err)
	}
	//由于c是[]byte 类型，fmt.Println 会以默认格式输出每个字节的十进制数值
	fmt.Println("\n", c)
	//将[]byte 转换为字符串
	fmt.Println("\n", string(c))
}
