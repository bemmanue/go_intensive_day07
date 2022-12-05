package main

import (
	"log"
	"testing"
)

func TestFirst(t *testing.T) {
	coins := []int{1, 5, 10, 50, 100, 1000}

	result := minCoins2(1864, coins)

	coinsMap := make(map[int]int)
	for _, coin := range coins {
		coinsMap[coin] = 0
	}

	for _, coin := range result {
		if _, ok := coinsMap[coin]; !ok {
			log.Fatalln("Test 1: result contains nonexistant coin")
		}
		coinsMap[coin]++
	}

	for i := 0; i+1 < len(coins); i++ {
		if coinsMap[coins[i]] >= coins[i+1]/coins[i] {
			log.Fatalln("Test 1: returned slice of coins not of minimal size")
		}
	}
}

func TestSecond(t *testing.T) {
	coins := []int{1, 5, 10, 1, 50, 5, 100, 10, 1000}

	result := minCoins2(1864, coins)

	coins = []int{1, 5, 10, 50, 100, 1000}

	coinsMap := make(map[int]int)
	for _, coin := range coins {
		coinsMap[coin] = 0
	}

	for _, coin := range result {
		if _, ok := coinsMap[coin]; !ok {
			log.Fatalln("Test 2: result contains nonexistant coin")
		}
		coinsMap[coin]++
	}

	for i := 0; i+1 < len(coins); i++ {
		if coinsMap[coins[i]] >= coins[i+1]/coins[i] {
			log.Fatalln("Test 2: returned slice of coins not of minimal size")
		}
	}
}

func BenchmarkMinCoins(b *testing.B) {
	coins := []int{1, 5, 10, 50, 100, 1000}
	for i := 0; i < b.N; i++ {
		minCoins2(1864, coins)
	}
}