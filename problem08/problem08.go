package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input08.txt
var problem08input string

func main() {
	first()
	second()
}

func parseOutputs() [][]string {
	var outputVals [][]string

	for _, line := range util.Lines(problem08input) {
		signal := strings.Split(line, "|")
		outputVals = append(outputVals, strings.Fields(signal[1]))
	}

	return outputVals
}

func sortOutputs(raw string) string {
	var out []string
	for _, c := range "abcdefg" {
		char := string(c)
		if strings.Contains(raw, char) {
			out = append(out, char)
		}
	}

	return strings.Join(out, "")
}

func parseLookupTable(raw string) map[string]string {
	allParts := strings.Fields(raw)

	lookup := make(map[string]string)
	counts := make(map[rune]int)

	for _, parts := range allParts {
		for _, item := range parts {
			counts[item]++
		}
	}

	for _, part := range allParts {
		fmt.Println()
		fmt.Println(part)
		sorted := sortOutputs(part)
		switch len(part) {
		case 2:
			lookup[sorted] = "1"
		case 3:
			lookup[sorted] = "7"
		case 4:
			lookup[sorted] = "4"
		case 7:
			lookup[sorted] = "8"
		case 6:
			innerCounts := make(map[int]int)
			for _, c := range part {
				innerCounts[counts[c]]++
			}

			fmt.Println(innerCounts)

			if innerCounts[7] == 1 {
				lookup[sortOutputs(part)] = "0"
			} else if innerCounts[4] == 0 {
				lookup[sortOutputs(part)] = "9"
			} else {
				lookup[sortOutputs(part)] = "6"
			}
		case 5:
			for _, c := range part {
				if counts[c] == 4 {
					lookup[sortOutputs(part)] = "2"
				} else if counts[c] == 6 {
					lookup[sortOutputs(part)] = "5"
				}
			}

			if _, ok := lookup[sortOutputs(part)]; !ok {
				lookup[sortOutputs(part)] = "3"
			}
		}
	}

	return lookup
}

func first() {
	var count int
	for _, outputSequence := range parseOutputs() {
		for _, outputVal := range outputSequence {
			switch len(outputVal) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	fmt.Println(count)
}

func second() {
	var total int
	for _, line := range util.Lines(problem08input) {
		signal := strings.Split(line, "|")
		lookup := parseLookupTable(signal[0])

		var final []string
		for _, part := range strings.Fields(signal[1]) {
			finalItem := lookup[sortOutputs(part)]
			final = append(final, finalItem)
		}

		finalOutput := strings.Join(final, "")
		total += util.FInt(finalOutput)
	}

	fmt.Println(total)
}
