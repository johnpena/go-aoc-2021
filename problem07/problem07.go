package main

import (
	_ "embed"
	"fmt"
	"math"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input07.txt
var problem07input string

func main() {
	first()
	second()
}

func first() {
	positions := util.Ints(problem07input)
	cheapest := util.Median(positions)
	totalCost := 0
	for _, i := range positions {
		cost := util.AbsInt(i - cheapest)
		totalCost += cost
	}

	fmt.Println(totalCost)
}

func fuelUsage(distance int) int {
	return (distance * (distance + 1)) / 2
}

func second() {
	positions := util.Ints(problem07input)
	min := util.Min(positions...)
	max := util.Max(positions...)

	minFuel := math.MaxInt64
	for endPosition := min; endPosition <= max; endPosition++ {
		amount := 0
		for _, startPosition := range positions {
			fuelUsed := fuelUsage(util.AbsInt(endPosition - startPosition))
			amount += fuelUsed
		}
		if amount < minFuel {
			minFuel = amount
		}
	}

	fmt.Println(minFuel)
}
