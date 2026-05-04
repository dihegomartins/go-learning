package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 5
	resultado := fizzBuzz(n)
	fmt.Println(resultado)
}

func fizzBuzz(n int) []string {
	answer := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}
	}
	return answer
}
