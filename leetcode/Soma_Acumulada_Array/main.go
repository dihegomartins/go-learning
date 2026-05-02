/*1480. Soma acumulada de um array unidimensional
LINK: https://leetcode.com/problems/running-sum-of-1d-array/description/
*/

package main

import "fmt"

func main() {
	dados := []int{3, 1, 2, 10, 1}

	resultado := runningSum(dados)

	fmt.Println(resultado)
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
