package main

import "fmt"

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("%+v\n", myArray)
	for index, _ := range myArray {
		if index == 2 {
			myArray[index] = "smart"
		} else if index == 4 {
			myArray[index] = "strong"
		}
	}
	fmt.Printf("%+v\n", myArray)
}
