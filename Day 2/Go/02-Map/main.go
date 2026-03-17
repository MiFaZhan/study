package main

import "fmt"

func main() {
	m := make(map[string]int)

	//Map 是引用类型，传递时不拷贝底层数据结构
	m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// 遍历map
	//打印结果可能是无序的（Go 的 Map 遍历顺序是随机的）
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}

	v1 := m["a"]
	fmt.Println("\n获取a的值:", v1)

	v2, ok := m["d"]
	fmt.Println("\n获取d的值:", v2, ok)

	m["d"] = 4
	fmt.Println("\n修改d的值:", m["d"])

	delete(m, "d")
	fmt.Println("\n删除d后map的长度:", len(m))
}

/*
Map 类型的 key-value 可以是多种数据类型组合，以下是一些常用的类型：

* **string-string**: 最常见的类型，例如 `map[string]string` 用于存储键值对，其中键和值都是字符串。
* **string-int**: 用于存储键为字符串，值为整数的键值对，例如 `map[string]int`。
* **string-struct**: 用于存储键为字符串，值为自定义结构体的键值对，例如 `map[string]MyStruct`。
* **string-slice**: 用于存储键为字符串，值为切片的键值对，例如 `map[string][]int`。
* **string-map**: 用于存储键为字符串，值为另一个 Map 的键值对，例如 `map[string]map[string]int`。
* **string-interface**: 用于存储键为字符串，值为接口的键值对，例如 `map[string]interface{}`，这种类型可以存储任何类型的数据。

此外，还可以根据需要创建其他类型的 key-value，例如：

* **int-string**: 类似于 string-int，但键为整数。
* **int-slice**: 类似于 string-slice，但键为整数。
* **int-map**: 类似于 string-map，但键为整数。
*/
