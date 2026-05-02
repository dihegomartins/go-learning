/* Dado um array de inteiros nums, mova todos 0os 's para o final do array, mantendo a ordem relativa dos elementos diferentes de zero.
Observe que você deve fazer isso no próprio local, sem criar uma cópia da matriz.
LINK: https://leetcode.com/problems/move-zeroes/description/
*/

package main

import "fmt"

func main() {

	meusDados := []int{0, 1, 0, 3, 12}

	resultado := moveZeroes(meusDados)
	fmt.Println(resultado)
}

func moveZeroes(nums []int) []int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for k < len(nums) {
		nums[k] = 0
		k++
	}

	return nums
}
