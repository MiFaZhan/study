package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. 使用 Reader 和 Writer
	r := strings.NewReader("Hello, Go!")
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	if err != nil {
		fmt.Println("copy error:", err)
		return
	}
	fmt.Printf("copied %d bytes: %s\n", n, buf.String())

	// 2. 使用 MultiWriter 同时写入文件和控制台
	file, _ := os.Create("Day 3/go/04-io/test.txt")
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	io.WriteString(mw, "This goes to both stdout and file.\n")

	// 3. 使用 LimitReader 限制读取长度
	limited := io.LimitReader(strings.NewReader("1234567890"), 5)
	data, _ := io.ReadAll(limited)
	fmt.Printf("limited read: %s\n", data) // 输出 "12345"

	// 4. 使用 TeeReader 同时读取并写入
	var teeBuf bytes.Buffer
	tee := io.TeeReader(strings.NewReader("tee example"), &teeBuf)
	io.Copy(os.Stdout, tee) // 输出到 stdout，同时 teeBuf 也会得到数据
	fmt.Printf("\ncaptured in buffer: %s\n", teeBuf.String())
}
