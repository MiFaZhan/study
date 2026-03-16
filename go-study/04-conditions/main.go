package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	switch age {
	case 18:
		fmt.Println("18岁")
	case 20:
		fmt.Println("20岁")
	default:
		fmt.Println("其他年龄")
	}

	// 类型判断
	typeSwitch(100)
	typeSwitch(3.14)
	typeSwitch(true)
	typeSwitch("字符串")
	typeSwitch([]int{1, 2, 3})
}

func typeSwitch(v any) {
	switch v.(type) {
	case int:
		fmt.Println("整型")
	case float64:
		fmt.Println("浮点数")
	case bool:
		fmt.Println("布尔值")
	case string:
		fmt.Println("字符串")
	default:
		fmt.Println("其他类型")
	}
}
