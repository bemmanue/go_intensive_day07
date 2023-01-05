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
	coins = putInOrder(coins)

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
				mincoins[coin] = memo[i-coin]
			}
		}

		m := minimum(mincoins)

		memo[i] = append(mincoins[m], m)
	}
	return memo[val]
}

func putInOrder(coins []int) []int {
	coins = removeDuplicates(coins)

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[i]
	})

	return coins
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
