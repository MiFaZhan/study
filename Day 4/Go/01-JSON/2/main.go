package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,string"`
}

func main() {
	//解析到结构体
	a := []byte(`{"name":"米法展","age":"18"}`)
	var p Person
	err := json.Unmarshal(a, &p)
	if err != nil {
		fmt.Printf("json err:%v\n", err)
		return
	}
	fmt.Println(p)

	//解析到接口
	b := []byte(`{"name":"枯藤","age":"18"}`)

	var i interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		fmt.Printf("json err:%v\n", err)
		return
	}
	fmt.Println(i)

	//json array解析到结构体数组
	jsonArr := []byte(`[{"name":"米法展","age":"18"},{"name":"枯藤","age":"20"}]`)
	var arr []Person
	err = json.Unmarshal(jsonArr, &arr)
	if err != nil {
		fmt.Printf("json err:%v\n", err)
		return
	}
	fmt.Println("\n", arr)

}
