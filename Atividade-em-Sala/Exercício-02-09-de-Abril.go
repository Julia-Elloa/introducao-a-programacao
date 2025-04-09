package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)

	m := maior(num1, num2, num3)

	fmt.Println(m)

}

func maior(num1, num2, num3 int) int {

	if num1 >= num2 && num1 >= num3 {
		return num1
	} else if num2 >= num3 {
		return num2
	} else {
		return num3
	}
}
