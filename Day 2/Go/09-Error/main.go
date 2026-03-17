package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func main() {
	err := errors.New("这是一个错误")
	fmt.Println(err)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
