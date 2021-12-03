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
	var transpose [12][1000]int

	lines := util.Lines(problem03input)

	for lineNum, line := range lines {
		for charNum, char := range line {
			i, err := strconv.Atoi(string(char))
			util.Check(err)
			transpose[charNum][lineNum] = i
		}
	}

	var mostCommonElement []int
	for _, array := range transpose {
		count1s := util.CountMatches(array[:], 1)
		count0s := util.CountMatches(array[:], 0)
		if count1s >= count0s {
			mostCommonElement = append(mostCommonElement, 1)
		} else {
			mostCommonElement = append(mostCommonElement, 0)
		}
	}


	o2Rating := search(lines, mostCommonElement, true)
	co2Rating := search(lines, mostCommonElement, false)

	fmt.Println(o2Rating)
	fmt.Println(co2Rating)
	fmt.Println(o2Rating * co2Rating)
}

func search(lines []string, mostCommon []int, filter bool) int {
	haystack := make([]string, len(lines))
	copy(haystack, lines)

	position := 0
	for len(haystack) > 1 {
		var itemsLeft []string
		mostCommonForPosition := mostCommon[position]
		for _, line := range haystack {
			char, err := strconv.Atoi(string(line[position]))
			util.Check(err)
			if (char == mostCommonForPosition) == filter {
				itemsLeft = append(itemsLeft, line)
			}
		}

		haystack = itemsLeft
		position++
	}


	item, err := strconv.ParseInt(haystack[0], 2, 64)
	util.Check(err)

	return int(item)
}

