package main

import (
	"fmt"
	"os"
)

func main() {
	//Print
	fmt.Println("Hello, World!")

	//Fprint
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./Day 3/go/02-Fmt/Fprint.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	n := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", n)

	//Sprint
	//fmt.Sprint 可以对输入参数进行格式化输出。例如，你可以将多个值格式化成一个字符串，或者将值转换为特定格式
	s1 := fmt.Sprint("年龄：", 18, "，名字：", "枯藤")
	fmt.Println(s1) // 输出：年龄：18，名字：枯藤

	//Errorf
	err = fmt.Errorf("这是一个错误")
	fmt.Println(err) // 输出：这是一个错误

	//格式化占位符
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	a := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("100%%\n")

	//fmt.Scan
	//Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
	var (
		name    string
		age     int
		married bool
	)

	//fmt.Scan(&name, &age, &married)
	//fmt.Printf("Scan结果 name:%s age:%d married:%t \n", name, age, married)

	//fmt.Scanf
	//Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中
	//fmt.Scanf("姓名：%s 年龄：%d 婚否：%t", &name, &age, &married)
	//fmt.Printf("Scanf结果 name:%s age:%d married:%t \n", name, age, married)

	//fmt.Scanln
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("\n扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

/*
占位符	说明
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
*/

/*
os.O_RDONLY	只读模式（默认）
os.O_WRONLY	只写模式
os.O_RDWR	读写模式
os.O_CREATE	文件不存在时创建
os.O_APPEND	追加模式（写入追加到末尾）
os.O_TRUNC	截断模式（打开时清空文件内容）
os.O_EXCL	排他模式（文件存在则报错）
os.O_SYNC	同步模式（写入时同步到磁盘）
*/
