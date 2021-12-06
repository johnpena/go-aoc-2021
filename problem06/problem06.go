package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input06.txt
var problem06input string

func main() {
	first()
	second()
}

func GrowthAfterDays(days int) int {
	input := strings.Split(problem06input, ",")

	var fishCountByDay [9]int
	for _, in := range input {
		daysRemaining := util.FInt(in)
		fishCountByDay[daysRemaining]++
	}

	for i := 0; i < days; i++ {
		newFish := fishCountByDay[0]
		for rotation := range fishCountByDay[:8] {
			fishCountByDay[rotation] = fishCountByDay[rotation+1]
		}

		fishCountByDay[8] = newFish
		fishCountByDay[6] += newFish
	}

	var count int
	for _, numFish := range fishCountByDay {
		count += numFish
	}

	return count
}

func first() {
	fmt.Println(GrowthAfterDays(80))
}

func second() {
	fmt.Println(GrowthAfterDays(256))
}
