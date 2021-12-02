package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input01.txt
var problem1input string

func main() {
	first()
	second()
}

func first() {
	increased := 0
	lastVal := -1
	for _, line := range util.Lines(problem1input) {
		val, err := strconv.Atoi(line)
		util.Check(err)

		if lastVal == -1 {
			lastVal = val
			continue
		}

		if val > lastVal {
			increased++
		}

		lastVal = val
	}
	fmt.Println(increased)
}

func second() {
	var windows []int
	inputs := util.Lines(problem1input)
	for i := 0; i < len(inputs) - 2; i++ {
		w1, err1 := strconv.Atoi(inputs[i])
		w2, err2 := strconv.Atoi(inputs[i+1])
		w3, err3 := strconv.Atoi(inputs[i+2])

		fmt.Println(w1, w2, w3)

		util.Check(err1, err2, err3)

		windows = append(windows, w1 + w2 + w3)
	}

	increased := 0
	lastWindow := -1
	for _, window := range windows {
		if lastWindow == -1 {
			lastWindow = window
			continue
		}

		if window > lastWindow {
			increased++
		}

		lastWindow = window
	}

	fmt.Println(increased)
}