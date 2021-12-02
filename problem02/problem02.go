package main

import (
	_ "embed"
	"fmt"
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
	for _, line := range util.Lines(problem02input) {
		parts := strings.Split(line, " ")
		fmt.Println(parts)
	}
}

func second() {
}
