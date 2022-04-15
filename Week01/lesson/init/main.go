package main

import (
	"fmt"
	_ "github.com/MiracleWong/CloudNativeCamp/Week01/lesson/init/a"
	_ "github.com/MiracleWong/CloudNativeCamp/Week01/lesson/init/b"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("Hello World")
}
