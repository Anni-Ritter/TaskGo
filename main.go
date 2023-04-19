package main

import (
	"fmt"
)

func main() {
	//task_1
	var num int
	fmt.Println("Введите число:")
	fmt.Scanln(&num)
	var digit1 int = num / 100
	var digit2 int = num/10 - digit1*10
	var digit3 int = num % 10
	var result int = digit1 + digit2*10 + digit3*100
	fmt.Println(result)

	//task_2
	var a int
	var b int
	var c int
	fmt.Println("Введите стороны a, b, c:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Не прямоугольный")
	}

	//task_3
	var k int
	var hours int
	var minutes int
	fmt.Scanln(&k)
	if 0 < k && k < 86399 {
		hours = k / 3600
		minutes = (k - hours*3600) / 60
		fmt.Printf("It is %d hours %d minutes", hours, minutes)
	} else {
		fmt.Println("Введено некорректное число")
	}
}
