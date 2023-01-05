package main

import (
	"log"
	"testing"
)

func TestFirst(t *testing.T) {
	coins := []int{1, 5, 10, 50, 100, 1000}
	result := minCoins(1864, coins)
	//log.Println(result)

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
	coins := []int{1000, 100, 50, 10, 5, 1}

	result := minCoins(1864, coins)
	log.Println(result)

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

func TestThird(t *testing.T) {
	coins := []int{1, 5, 10, 1, 50, 5, 100, 10, 1000}

	result := minCoins(1864, coins)
	log.Println(result)

	coins = []int{1, 5, 10, 50, 100, 1000}

	coinsMap := make(map[int]int)
	for _, coin := range coins {
		coinsMap[coin] = 0
	}

	for _, coin := range result {
		if _, ok := coinsMap[coin]; !ok {
			log.Fatalln("Test 3: result contains nonexistant coin")
		}
		coinsMap[coin]++
	}

	for i := 0; i+1 < len(coins); i++ {
		if coinsMap[coins[i]] >= coins[i+1]/coins[i] {
			log.Fatalln("Test 3: returned slice of coins not of minimal size")
		}
	}
}
