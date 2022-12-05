package main

import (
	"sort"
	"testing"
)

func TestFirst(t *testing.T) {
	coins := []int{1, 5, 10, 50, 100, 1000}
	result := minCoins(186, coins)

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
}
