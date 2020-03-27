package main

import "fmt"

func main() {
	header := `
	***************************
	Counting odd numbers
	***************************
	`
	fmt.Println(header)

	var firstNumber int
	fmt.Println("Type the first number:")
	fmt.Scanln(&firstNumber)

	var secondNumber int
	fmt.Println("Type the second number:")
	fmt.Scanln(&secondNumber)

	var countOddNums int

	for i := firstNumber; i <= secondNumber; i++ {
		if i%2 != 0 {
			countOddNums++
			fmt.Printf("The number \"%d\" is an odd one\n", i)
		}
	}

	header = `
	****************************
	Odd numbers counter's result
	****************************
	`
	fmt.Println(header)

	fmt.Printf("Between the numbers \"%d\" and \"%d\", there are %d odd numbers.", firstNumber, secondNumber, countOddNums)
}
