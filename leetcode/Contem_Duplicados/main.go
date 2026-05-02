package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}

	resultado := containsDuplicate(nums)
	fmt.Println(resultado)
}

func containsDuplicate(nums []int) bool {
	estoque := map[int]bool{}

	for _, num := range nums {
		if _, ok := estoque[num]; ok {
			return true
		}
		estoque[num] = true
	}
	return false
}
