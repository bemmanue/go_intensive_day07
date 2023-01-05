// My package
package main

import "sort"

// My function
// function returns minimum coins
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

func removeDuplicates(array []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}

	for _, item := range array {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
