package main

import "fmt"

func main() {
	hashmap := make(map[string]int)

	hashmap["a"] = 1
	hashmap["b"] = 2

	value, exists := hashmap["a"]
	if exists {
		fmt.Printf("a的值为: %d\n", value)
	}
	fmt.Printf("当前map: %v\n", hashmap)

	delete(hashmap, "a")
	fmt.Printf("删除a后: %v\n", hashmap)

	fmt.Println("遍历map:")
	for key, value := range hashmap {
		fmt.Printf("  key: %s, value: %d\n", key, value)
	}

	if _, ok := hashmap["a"]; ok {
		fmt.Println("a存在")
	}

}
