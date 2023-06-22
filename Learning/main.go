package main

import "fmt"

//var price = 100

func main() {

	msg := "Hello Ball's World"
	msg = "Thanachart Saejueng" + " Haha"
	//var age int = 23
	//var price float64 = 34.1414
	//var check bool = true

	//age := 23
	//price := 34.1414
	//check := true

	//var age, price, check = 23, 34.1414, true

	age, price, check := 23, 34.1414, true

	fmt.Println("msg:", msg)
	fmt.Println("age: ", age)
	fmt.Println("price: ", price)
	fmt.Println("check: ", check)
}
