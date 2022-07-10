package main

import (
	"flag"
	"fmt"
	"github.com/MiracleWong/CloudNativeCamp/Week10/homework/httpserver/metrics"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	flag.Set("v", "4")
	glog.V(2).Info("Starting http server ...")
	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.Handle("/metrics", promhttp.Handler())
	// 加入Debug模块
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal("start server failed: %s \n", err.Error())
	}

	// 监控两个信号
	// TERM信号（kill + 进程号 触发）
	// 中断信号（ctrl + c 触发）
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	s := <-osc
	fmt.Println("监听到退出信号,s=", s)

	// 退出前的清理操作
	// clean()

	fmt.Println("main程序退出")
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

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(4).Info("entering root handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	user := r.URL.Query().Get("user")
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello"+user))
	} else {
		io.WriteString(w, fmt.Sprintf("hello stranger "))
	}
	io.WriteString(w, "============Details of the http request header ==============")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf(k, v))
	}
	glog.V(4).Info("Respond in %d ms", delay)
}
