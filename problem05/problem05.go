package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/johnpena/go-aoc-2021/util"
)

//go:embed input05.txt
var problem05input string

func main() {
	first()
	second()
}

type Line struct {
	X1, Y1, X2, Y2 int
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.X1, l.Y1, l.X2, l.Y2)
}

func (l *Line) ParseFrom(str string) {
	parts := strings.Split(str, " -> ")
	firstPart := strings.Split(parts[0], ",")
	secondPart := strings.Split(parts[1], ",")
	l.X1, _ = strconv.Atoi(firstPart[0])
	l.Y1, _ = strconv.Atoi(firstPart[1])
	l.X2, _ = strconv.Atoi(secondPart[0])
	l.Y2, _ = strconv.Atoi(secondPart[1])
}

func (l *Line) IsVertical() bool {
	return l.X1 == l.X2
}

func (l *Line) IsHorizontal() bool {
	return l.Y1 == l.Y2
}

func (l *Line) IsDiagonal() bool {
	return !(l.IsVertical() || l.IsHorizontal())
}

type Grid struct {
	ConsiderDiagonal bool
	cells            [999][999]int
}

func (grid *Grid) Print() {
	for _, row := range grid.cells {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (grid *Grid) AddLine(l Line) {
	if l.IsDiagonal() {
		if !grid.ConsiderDiagonal {
			return
		}

		x := l.X1
		xEnd := l.X2
		xDirection := 1
		if x > xEnd {
			xDirection = -1
		}

		y := l.Y1
		yEnd := l.Y2
		yDirection := 1
		if y > yEnd {
			yDirection = -1
		}

		for {
			grid.cells[x][y]++
			x += xDirection
			y += yDirection

			if x == xEnd || y == yEnd {
				grid.cells[x][y]++
				x += xDirection
				y += yDirection
				break
			}
		}
	} else if l.IsHorizontal() {
		x := util.Min(l.X1, l.X2)
		xEnd := util.Max(l.X1, l.X2) + 1
		for i := x; i < xEnd; i++ {
			grid.cells[i][l.Y1]++
		}
	} else {
		y := util.Min(l.Y1, l.Y2)
		yEnd := util.Max(l.Y1, l.Y2) + 1
		for i := y; i < yEnd; i++ {
			grid.cells[l.X1][i]++
		}
	}
}

func (grid *Grid) CountOverlaps() int {
	count := 0
	for _, row := range grid.cells {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}

	return count
}

func first() {
	grid := &Grid{
		ConsiderDiagonal: false,
	}
	for _, line := range util.Lines(problem05input) {
		var l Line
		l.ParseFrom(line)
		grid.AddLine(l)
	}

	fmt.Println(grid.CountOverlaps())
}

func second() {
	grid := &Grid{
		ConsiderDiagonal: true,
	}
	for _, line := range util.Lines(problem05input) {
		var l Line
		l.ParseFrom(line)
		grid.AddLine(l)
	}

	fmt.Println(grid.CountOverlaps())
}
