package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 解析 URL 参数
	data := r.URL.Query()
	name := data.Get("name")
	age := data.Get("age")

	fmt.Printf("收到请求 - name: %s, age: %s\n", name, age)

	// 返回 JSON 响应
	answer := `{"status": "ok", "message": "请求成功"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(answer))
}

func main() {
	// 注册路由
	http.HandleFunc("/get", getHandler)

	// 启动服务器
	addr := "127.0.0.1:9090"
	fmt.Printf("服务器正在监听 %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("服务器启动失败：%v\n", err)
	}
}
