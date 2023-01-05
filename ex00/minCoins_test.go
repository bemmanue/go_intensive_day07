package main

import (
	"fmt"
	"strconv"
	"testing"
)

var value int = 13

// Паттерн "Табличные тесты"
func TestMinCoins(t *testing.T) {
	coinsSet := [][]int{
		{1, 5, 10, 50, 100, 500, 1000},
		{1, 3, 4, 7, 13, 15},
	}

	for i, coins := range coinsSet {
		t.Run("Test_"+strconv.Itoa(i+1), func(t *testing.T) {
			result := minCoins2(value, coins)
			//result := minCoins2(value, coins)
			fmt.Println(result)

			//coinsMap := make(map[int]int)
			//for _, coin := range coins {
			//	coinsMap[coin] = 0
			//}
			//
			//for _, coin := range result {
			//	if _, ok := coinsMap[coin]; !ok {
			//		t.Error("result contains nonexistant coin")
			//	}
			//	coinsMap[coin]++
			//}

			//for i := len(coins) - 1; i > 0; i-- {
			//	if coinsMap[coins[i]] == 0 {
			//		continue
			//	}
			//	if coinsMap[coins[i-1]] >= coins[i]/coins[i-1] {
			//		t.Error("returned slice of coins not of minimal size")
			//	}
			//}

			//for i := 0; i+1 < len(coins); i++ {
			//	if coinsMap[coins[i]] == 0 {
			//		continue
			//	}
			//	if coinsMap[coins[i]] >= coins[i+1]/coins[i] {
			//		t.Error("returned slice of coins not of minimal size")
			//	}
			//}
		})
	}
}
