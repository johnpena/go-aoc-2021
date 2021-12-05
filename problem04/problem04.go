package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input04.txt
var problem04input string

func main() {
	first()
	second()
}

type Board struct {
	cells   [5][5]int
	matches [5][5]bool
	wonAlready bool
}

func (b *Board) ParseFrom(s string) {
	for i, line := range util.Lines(s) {
		for j, cell := range strings.Fields(line) {
			cellVal, err := strconv.Atoi(cell)
			util.Check(err)
			b.cells[i][j] = cellVal
		}
	}
}

func (b *Board) SumUnmatched() int {
	var sum int
	for i, row := range b.matches {
		for j, cell := range row {
			if !cell {
				sum += b.cells[i][j]
			}
		}
	}

	return sum
}

func (b *Board) HasWon() bool {
	if b.wonAlready {
		return false
	}

	for _, row := range b.matches {
		nMatches := 0
		for _, match := range row {
			if match {
				nMatches++
			}
		}

		if nMatches == 5 {
			b.wonAlready = true
			return true
		}
	}

	for i := 0; i < 5; i++ {
		nMatches := 0
		for _, row := range b.matches {
			match := row[i]
			if match {
				nMatches++
			}
		}
		if nMatches == 5 {
			b.wonAlready = true
			return true
		}
	}

	return false
}

func (b *Board) SetLastBallDrawn(ball int) int {
	for i, row := range b.cells {
		for j, cell := range row {
			if cell == ball {
				b.matches[i][j] = true
			}
		}
	}

	if b.HasWon() {
		return b.SumUnmatched() * ball
	}

	return 0
}

func GetBallsAndBoards() ([]int, []*Board) {
	parts := strings.Split(problem04input, "\n\n")
	balls := parts[0]

	var boards []*Board
	for _, part := range parts[1:]{
		var board Board
		board.ParseFrom(part)
		boards = append(boards, &board)
	}

	var ballVals []int
	for _, ball := range strings.Split(balls, ",") {
		ballVal, err := strconv.Atoi(ball)
		util.Check(err)
		ballVals = append(ballVals,ballVal)
	}

	return ballVals, boards
}

func first() {
	balls, boards := GetBallsAndBoards()
	OUTER:
	for _, ball := range balls {
		for _, board := range boards {
			result := board.SetLastBallDrawn(ball)
			if result > 0 {
				fmt.Println(fmt.Sprintf("winner: %d", result))
				break OUTER
			}
		}
	}
}

func second() {
	balls, boards := GetBallsAndBoards()

	var winners int
OUTER:
	for _, ball := range balls {
		for _, board := range boards {
			result := board.SetLastBallDrawn(ball)
			if result > 0 {
				winners++
				if winners == len(boards){
					fmt.Println(fmt.Sprintf("winner: %d", result))
					break OUTER
				}
			}
		}
	}
}
