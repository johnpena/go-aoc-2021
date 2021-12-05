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

func transpose() [12][1000]int {
	var t [12][1000]int
	for lineNum, line := range util.Lines(problem03input) {
		for charNum, char := range line {
			i, err := strconv.Atoi(string(char))
			util.Check(err)
			t[charNum][lineNum] = i
		}
	}

	return t
}

func first() {
	t := transpose()

	var gammaElements []int
	var epsilonElements []int
	for _, array := range t {
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
	lines := util.Lines(problem03input)

	o2Rating := search(lines, 0, true)
	co2Rating := search(lines, 0, false)

	fmt.Println(o2Rating * co2Rating)
}

func search(haystack []string, criteria int, filter bool) int {
	if len(haystack) == 1 {
		needle, err := strconv.ParseInt(haystack[0], 2, 46)
		util.Check(err)
		return int(needle)
	}

	var zeroItems []string
	var oneItems []string
	for _, item := range haystack {
		if rune(item[criteria]) == '0' {
			zeroItems = append(zeroItems, item)
		} else {
			oneItems = append(oneItems, item)
		}
	}

	if (len(oneItems) >= len(zeroItems)) == filter {
		return search(oneItems, criteria+1, filter)
	}

	return search(zeroItems, criteria+1, filter)
}
