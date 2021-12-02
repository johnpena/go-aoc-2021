package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input02.txt
var problem02input string

func main() {
	first()
	second()
}

func first() {
	position := 0
	depth := 0
	for _, line := range util.Lines(problem02input) {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, err := strconv.Atoi(parts[1])
		util.Check(err)

		switch direction {
		case "forward":
			position += magnitude
		case "down":
			depth += magnitude
		case "up":
			depth -= magnitude
		default:
			panic(fmt.Sprintf("bad direction %s", direction))
		}
	}

	fmt.Println(position * depth)
}

func second() {
	position := 0
	depth := 0
	aim := 0
	for _, line := range util.Lines(problem02input) {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, err := strconv.Atoi(parts[1])
		util.Check(err)

		switch direction {
		case "forward":
			position += magnitude
			depth += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		default:
			panic(fmt.Sprintf("bad direction %s", direction))
		}
	}

	fmt.Println(position * depth)
}
