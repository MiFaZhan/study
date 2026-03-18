package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//file, err := os.Open("file.txt") // 用于读取访问。
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close() // 确保文件在使用完毕后关闭

	var home string
	switch runtime.GOOS {
	case "windows":
		home = os.Getenv("USERPROFILE")
	case "linux":
		home = os.Getenv("HOME")
	default:
		home = os.Getenv("HOME")
	}
	fmt.Println("HOME:", home)

	//硬编码
	//var ProjectPath string
	//ProjectPath = "Day 3/go/03-Os"

	//环境变量注入
	os.Setenv("ProjectPath", "Day 3/go/03-Os")

	//读取环境变量
	ProjectPath := os.Getenv("ProjectPath")
	fmt.Println("ProjectPath:", ProjectPath)

	//创建目录
	var err error
	err = os.Mkdir(ProjectPath+"/Test-01", 0755)
	if err != nil {
		fmt.Println("创建目录失败，err:", err)
		os.Exit(0)
	}

	// 创建文件并写入
	file, err := os.Create(ProjectPath + "/Test-01/file.txt")
	if err != nil {
		fmt.Println("创建文件失败，Create err:", err)
		os.Exit(1)
	}
	defer file.Close() // 确保文件在使用完毕后关闭

	// 写入内容
	file.WriteString("这是一个测试文件\n")

	//重命名文件
	file.Close() // 立即关闭
	os.Rename(ProjectPath+"/Test-01/file.txt", ProjectPath+"/Test-01/file-01.txt")

	//读取文件信息
	info, err := os.Stat(ProjectPath + "/Test-01/file-01.txt")
	if err != nil {
		fmt.Println("读取文件信息失败，err:", err)
		os.Exit(1)
	}
	fmt.Println("===== 文件信息 =====")
	fmt.Printf("文件名: %s\n", info.Name())
	fmt.Printf("文件大小: %d 字节\n", info.Size())
	fmt.Printf("修改时间: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("是否为目录: %v\n", info.IsDir())
	fmt.Printf("权限模式: %s\n", info.Mode())
}
