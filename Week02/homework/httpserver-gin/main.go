package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		// 1. 接收客户端 request，并将 request 中带的 header 写入 response header——未实现
		fmt.Println("Request Header: ", c.Request.Header)
		//for k, v := range c.Request.Header {
		//	fmt.Println(k, v)
		//}
		//fmt.Println("Response Header: ", c.Writer.Header().Clone())

		// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		// 塞入一个默认的VERSION值，防止测试用例为空
		os.Setenv("VERSION", "go1.17.2")

		version := os.Getenv("VERSION")
		fmt.Println("VERSION is:", version)
		c.Writer.Header().Add("VERSION", version)
		fmt.Println("Response Header: ", c.Writer.Header())

		// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		ip, err := c.RemoteIP()
		if err != true {
			fmt.Println("err:", err)
		}

		if net.IP(ip) != nil {
			fmt.Println("Client ip: ", ip)
			log.Println(ip)
		}

		fmt.Println("http Status Code: ", http.StatusOK)
		log.Println(http.StatusOK)

		// 4. 当访问 localhost/healthz 时，应返回 200
		c.Writer.WriteString(strconv.Itoa(http.StatusOK))
	})

	r.Run(":80") // listen and serve on 0.0.0.0:80
}
