/*É fornecido um array pricesonde prices[i]`price` representa o preço de uma determinada ação no dia.ith
O objetivo é maximizar o lucro escolhendo um único dia para comprar uma ação e outro dia, em uma data futura diferente, para vendê-la.
Retorne o lucro máximo que você puder obter com esta transação . Se você não puder obter nenhum lucro, devolva o valor 0.
LINK: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
*/

package main

import "fmt"

func main() {
	meusDados := []int{7, 6, 4, 3, 1}
	resultado := maxProfit(meusDados)
	fmt.Println(resultado)
}

func maxProfit(prices []int) int {
	left := 0
	max_profit := 0
	for r := 1; r < len(prices); r++ {
		profift := prices[r] - prices[left]
		max_profit = max(profift, max_profit)
		if profift < 0 {
			left = r
		}
	}
	return max_profit
}
