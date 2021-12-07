package util

import (
	"math"
	"sort"
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

func Max(items ...int) int {
	max := 0
	for _, i := range items {
		if i > max {
			max = i
		}
	}

	return max
}

func Min(items ...int) int {
	min := math.MaxInt64
	for _, i := range items {
		if i < min {
			min = i
		}
	}

	return min
}

func Mean(items []int) float64 {
	size := len(items)
	if size == 0 {
		panic("can't take a mean of zero length array")
	}

	sum := 0
	for _, item := range items {
		sum += item
	}

	return float64(sum) / float64(size)
}

func Median(items []int) int {
	copiedItems := make([]int, len(items))
	copy(copiedItems, items)
	sort.Ints(copiedItems)

	size := len(copiedItems)

	if size == 0 {
		panic("can't take a median of zero length array")
	} else if size % 2 == 0 {
		return int(Mean(copiedItems[size/2-1 : size/2+1]))
	} else {
		return copiedItems[size/2]
	}
}

func FInt(str string) int {
	x, err := strconv.Atoi(str)
	Check(err)

	return x
}

func AbsInt(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func Ints(input string) []int {
	var ints []int
	strs := strings.Split(input, ",")
	for _, i := range strs {
		ints = append(ints, FInt(i))
	}

	return ints
}