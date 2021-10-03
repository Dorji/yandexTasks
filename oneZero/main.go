package main

import (
	"fmt"
	"strings"
)

func main() {
	var first string
	var second string

	var firstSlice []string
	var secondSlice []string

	var firstNum string
	var secondNum string

	fmt.Scanln(&first)
	fmt.Scanln(&second)

	firstSlice = strings.Split(first, "")
	secondSlice = strings.Split(second, "")

	wait := 0
	for k, v := range firstSlice {
		if k > wait || wait == 0 {

			if v == "z" {
				firstNum = firstNum + "0"
				wait = k + 3
			}
			if v == "o" {
				firstNum = firstNum + "1"
				wait = k + 2
			}
		}

		if k == wait {
			wait = 0
		}
	}

	wait = 0
	for k, v := range secondSlice {
		if k > wait || wait == 0 {
			if v == "z" {
				secondNum = secondNum + "0"
				wait = k + 3
			}
			if v == "o" {
				secondNum = secondNum + "1"
				wait = k + 2
			}
		}
		if k == wait {
			wait = 0
		}
	}
	a := 0
	if len(firstNum) > len(secondNum) {
		fmt.Println(">")
	} else if len(firstNum) < len(secondNum) {
		fmt.Println("<")
	} else {

		for k := range firstNum {
			if a == 0 {
				if firstNum[k] > secondNum[k] {
					fmt.Println(">")
					a = 1
				} else {
					fmt.Println("<")
					a = 1
				}
			}
		}
		if a == 0 {
			fmt.Println("=")
		}
	}
}
