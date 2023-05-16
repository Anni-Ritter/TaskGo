package main

import (
	"fmt"
	"strconv"
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

	//task_4
	var str string
	var newStr string
	fmt.Println("Введите строку не больше 1000 символов: ")
	fmt.Scanln(&str)
	if len(str) > 1000 {
		fmt.Println("Не больше 1000 символов")
	} else {
		for i := 0; i < len(str); i++ {
			if i != 0 {
				newStr += "*"
			}
			newStr += string(str[i])
		}
		fmt.Println(newStr)
	}

	//task_5
	var number int
	fmt.Println("Введите число:")
	fmt.Scanln(&number)
	strNum := strconv.Itoa(number)
	newNum := ""
	for i := 0; i < len(strNum); i++ {
		digit, _ := strconv.Atoi(string(strNum[i]))
		square := digit * digit
		newNum += strconv.Itoa(square)
	}
	results, _ := strconv.Atoi(newNum)
	fmt.Println(results)
}
