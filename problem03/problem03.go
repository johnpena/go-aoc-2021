package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input03.txt
var problem03input string

func main() {
	first()
	second()
}

func decode(binary []int) int {
	var s string
	for _, item := range binary {
		s += strconv.Itoa(item)
	}

	result, err := strconv.ParseInt(s, 2, 64)
	util.Check(err)

	return int(result)
}

func first() {
	var transpose [12][1000]int

	for lineNum, line := range util.Lines(problem03input) {
		for charNum, char := range line {
			i, err := strconv.Atoi(string(char))
			util.Check(err)
			transpose[charNum][lineNum] = i
		}
	}

	var gammaElements []int
	var epsilonElements []int
	for _, array := range transpose {
		gammaElement := util.MostCommonElement(array[:])
		var epsilonElement int
		if gammaElement == 0 {
			epsilonElement = 1
		}

		gammaElements = append(gammaElements, gammaElement)
		epsilonElements = append(epsilonElements, epsilonElement)
	}

	gamma := decode(gammaElements)
	epsilon := decode(epsilonElements)

	fmt.Println(gamma * epsilon)
}

func second() {
}

