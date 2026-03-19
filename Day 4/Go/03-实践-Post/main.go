package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 演示 GET 请求
	if err := sendGet(); err != nil {
		fmt.Fprintf(os.Stderr, "GET 请求失败: %v\n", err)
	}

	fmt.Println()

	// 演示 POST 请求
	if err := sendPost(); err != nil {
		fmt.Fprintf(os.Stderr, "POST 请求失败: %v\n", err)
	}
}

// sendGet 发送 GET 请求，并在请求头中传递自定义参数，然后读取响应头
func sendGet() error {
	// 1. 创建请求
	url := "https://httpbin.org/get"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 2. 在请求头中设置参数（例如：自定义 User-Agent 和 X-Request-ID）
	req.Header.Set("User-Agent", "MyGoApp/1.0")
	req.Header.Set("X-Request-ID", "12345-abcde")
	// 还可以添加其他自定义头
	req.Header.Set("Accept", "application/json")

	// 3. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 4. 读取响应头中的指定参数值
	contentType := resp.Header.Get("Content-Type")
	server := resp.Header.Get("Server")
	date := resp.Header.Get("Date")
	// 如果有自定义响应头（本例中 httpbin 不会返回自定义头，但可以读取标准头）
	fmt.Printf("GET 响应头信息:\n")
	fmt.Printf("  Content-Type: %s\n", contentType)
	fmt.Printf("  Server: %s\n", server)
	fmt.Printf("  Date: %s\n", date)

	// 可选：读取响应体（仅用于查看返回内容，不是必须）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %w", err)
	}
	fmt.Printf("  Response Body (first 200 chars): %.200s...\n", body)

	return nil
}

// sendPost 发送 POST 请求，并在请求头中传递参数，然后读取响应头
func sendPost() error {
	// 1. 准备要发送的 JSON 数据（模拟表单或 API 参数）
	postData := map[string]interface{}{
		"name":  "Alice",
		"age":   30,
		"email": "alice@example.com",
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		return fmt.Errorf("JSON 编码失败: %w", err)
	}

	// 2. 创建 POST 请求
	url := "https://httpbin.org/post"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 3. 设置请求头（包含内容类型及自定义参数）
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "MyGoApp/1.0")
	req.Header.Set("X-Custom-Token", "s3cr3t-t0k3n")
	req.Header.Set("X-Request-ID", "67890-fghij")

	// 4. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 5. 读取响应头中的参数值
	contentType := resp.Header.Get("Content-Type")
	contentLength := resp.Header.Get("Content-Length")
	date := resp.Header.Get("Date")
	fmt.Printf("POST 响应头信息:\n")
	fmt.Printf("  Content-Type: %s\n", contentType)
	fmt.Printf("  Content-Length: %s\n", contentLength)
	fmt.Printf("  Date: %s\n", date)

	// 可选：读取响应体（httpbin.org/post 会将我们发送的数据回显）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %w", err)
	}
	fmt.Printf("  Response Body (first 200 chars): %.200s...\n", body)

	return nil
}
