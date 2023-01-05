package main

import (
	"sort"
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

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[i]
	})

	var memo = make([][]int, val+1)

	memo[0] = []int{}

	for _, coin := range coins {
		if coin <= val {
			memo[coin] = []int{coin}
		}
		if coin == val {
			return memo[coin]
		}
	}

	for i := 1; i <= val; i++ {
		for _, coin := range coins {
			if i == coin {
				continue
			}
		}

		var mincoins = make([][]int, len(coins))

		for j, coin := range coins {
			if i > coin {
				mincoins[j] = memo[i-coin]
			}
		}
		//
		m, k := min(mincoins)
		memo[i] = append(m, coins[k])
	}
	return memo[val]
}

func min(mincoins [][]int) ([]int, int) {
	var m, i int

	for i = range mincoins {
		if len(mincoins[i]) != 0 {
			m = i
			break
		}
	}

	for i = range mincoins {
		if len(mincoins[i]) == 0 {
			continue
		}
		if len(mincoins[m]) < len(mincoins[i]) {
			m = i
		}
	}

	return mincoins[m], m
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
