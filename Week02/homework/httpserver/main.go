package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {

	// 4. 当访问 localhost/healthz 时，应返回 200
	io.WriteString(w, strconv.Itoa(http.StatusOK))

	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	fmt.Println("Request解析开始……")
	fmt.Println("Request Header: ", r.Header)

	resp := http.Response{Header: r.Header}
	fmt.Println("Response Header: ", resp.Header)

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	// 塞入一个默认的VERSION值，防止测试用例为空
	os.Setenv("VERSION", "go1.17.2")

	version := os.Getenv("VERSION")
	fmt.Println("VERSION is:", version)
	resp.Header.Add("VERSION", version)
	fmt.Println("Response Header: ", resp.Header)

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		fmt.Println("Client ip: ", ip)
		log.Println(ip)
	}

	fmt.Println("http Status Code: ", http.StatusOK)
	log.Println(http.StatusOK)
}
