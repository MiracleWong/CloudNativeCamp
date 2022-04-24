package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	io.WriteString(w, strconv.Itoa(http.StatusOK))

	//HTTP方法
	fmt.Println("Request解析开始……")
	fmt.Println("Request Header: ", r.Header)

	resp := http.Response{Header: r.Header}
	fmt.Println("Response Header: ", resp.Header)

}
