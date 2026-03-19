package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//搜索关键字
	keyword := "go"
	if len(os.Args) > 1 {
		keyword = os.Args[1]
	}

	//构建搜索URL
	baseUrl := "https://cn.bing.com/search"
	params := url.Values{}
	params.Add("q", keyword)
	fullUrl := baseUrl + "?" + params.Encode()

	//创建HTTP客户端（可使用默认客户端）
	client := &http.Client{}

	//创建HTTP请求
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//设置User-Agent 模拟浏览器，避免被拒绝
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	//检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("请求失败", resp.StatusCode)
		return
	}

	//读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//打印响应体
	fmt.Println(string(body))
}
