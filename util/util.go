package util

import (
	"strconv"
	"strings"
)

func Check(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func MostCommonElement(array []int) int {
	elementCount := map[int]int{}
	var currentMax int
	var currentMostCommon int
	for _, item := range array {
		elementCount[item]++
		if elementCount[item] > currentMax {
			currentMax = elementCount[item]
			currentMostCommon = item
		}
	}

	return currentMostCommon
}

func CountMatches(array []int, filterValue int) int {
	var count int
	for _, item := range array {
		if item == filterValue {
			count++
		}
	}

	return count
}

func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func FInt(str string) int {
	x, err := strconv.Atoi(str)
	Check(err)

	return x
}