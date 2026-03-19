package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	apiUrl := "http://127.0.0.1:9090/get" // 定义请求的 API 地址
	// 构造 URL 查询参数
	data := url.Values{}
	data.Set("name", "枯藤") // 设置 name 参数
	data.Set("age", "18")  // 设置 age 参数

	// 解析基础 URL
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
		return // 解析失败则直接返回
	}
	// 将编码后的参数拼接到 URL 后面
	u.RawQuery = data.Encode()
	fmt.Println(u.String()) // 打印完整的请求 URL

	// 发送 GET 请求
	var resp *http.Response
	resp, err = http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err) // 修正日志描述为 get
		return
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 读取响应内容
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp failed,err:%v\n", err)
		return
	}
	// 打印响应结果
	fmt.Println(string(b))
}
