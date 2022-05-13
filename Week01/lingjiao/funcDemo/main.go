package main

import "fmt"

type animal struct {
	Name  string
	Color string
	Age   int
}

func (a *animal) Run() {
	fmt.Println(a.Name, "在奔跑")
}
func (a *animal) jump() {
	fmt.Println(a.Name, "跳")
}

type Cat struct {
	Cat animal
}

func main() {
	var cat1 = Cat{
		Cat: animal{
			Name:  "Tom",
			Color: "Black",
			Age:   10,
		},
	}
	fmt.Println(cat1.Cat.Name)
	cat1.Cat.Run()
}
