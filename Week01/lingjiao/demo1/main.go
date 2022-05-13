//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan int, 5)
//	go produce(ch)
//	consumer(ch)
//}
//
//func produce(ch chan<- int) {
//	for i := 0; i < 10; i++ {
//		ch <- i
//		time.Sleep(1 * time.Second)
//		fmt.Printf("Producing Data: %d\n", i)
//	}
//	close(ch)
//}
//
//func consumer(ch <-chan int) {
//	for k := range ch {
//		fmt.Printf("Get Data: %d\n", k)
//	}
//}
package main

import "fmt"

func main() {

	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(1, 2)
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)
}
