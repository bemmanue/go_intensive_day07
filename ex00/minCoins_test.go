package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

var value int = 6

// Паттерн "Табличные тесты"
func TestMinCoins(t *testing.T) {
	coinsSet := [][]int{
		{1, 5, 10, 50, 100, 500, 1000},
		{1, 3, 4, 7, 13, 15},
		{1000, 5, 10, 50, 100, 500, 1},
		{15, 13, 7, 4, 3, 1},
	}

	for i, coins := range coinsSet {
		t.Run("Test_"+strconv.Itoa(i+1), func(t *testing.T) {
			result := minCoins(value, coins)
			fmt.Println(result)

			coinsMap := make(map[int]bool)
			for _, coin := range coins {
				coinsMap[coin] = true
			}

			var sum int
			for _, coin := range result {
				if _, ok := coinsMap[coin]; !ok {
					t.Error("result contains nonexistant coin")
				}
				sum += coin
			}

			if len(result) != minCount(value, coins) {
				t.Error("returned slice of coins not of minimal size")
			}

			if len(result) != 0 && sum != value {
				t.Error("the sum of returned coins is not equal to expected value")
			}
		})
	}
}

func minCount(val int, coins []int) int {
	coins = putInOrder(coins)

	var max = math.MaxInt
	var memo = make([]int, val+1)

	for i := range memo {
		memo[i] = max
	}

	for _, coin := range coins {
		if val > coin {
			memo[coin] = 1
		}
	}

outer:
	for i := 2; i <= val; i++ {

		for _, coin := range coins {
			if i == coin {
				continue outer
			}
		}

		var mincoins = make([]int, len(coins))

		for j, coin := range coins {
			if i > coin {
				mincoins[j] = memo[i-coin]
			} else {
				mincoins[j] = max
			}
		}

		min := max

		for j := range mincoins {
			if mincoins[j] < min {
				min = mincoins[j]
			}
		}

		memo[i] = min + 1
	}
	return memo[val]
}
