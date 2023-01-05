package main

import (
	"net/http"
	_ "net/http/pprof"
)

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	coins = removeDuplicates(coins)

	var memo = make([][]int, val+1)

	for _, coin := range coins {
		if coin <= val {
			memo[coin] = []int{coin}
		}
		if coin == val {
			return memo[coin]
		}
	}

outer:
	for i := 2; i <= val; i++ {
		for _, coin := range coins {
			if i == coin {
				continue outer
			}
		}

		var mincoins = make(map[int][]int)

		for _, coin := range coins {
			if i > coin {
				mincoins[coin] = append(memo[i-coin], coin)
			}
		}

		m := minimum(mincoins)

		minSet := make([]int, len(mincoins[m]))
		copy(minSet, mincoins[m])

		memo[i] = minSet
	}
	return memo[val]
}

func removeDuplicates(array []int) []int {
	allKeys := make(map[int]bool)
	var list []int

	for _, item := range array {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func minimum(mincoins map[int][]int) int {
	var min int

	for i := range mincoins {
		min = i
		break
	}

	for m := range mincoins {
		if len(mincoins[m]) < len(mincoins[min]) {
			min = m
		}
	}

	return min
}

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
