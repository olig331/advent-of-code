package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	grid  [][]string
	folds []Fold
	yLen,
	xLen int
}

type Coords struct {
	x,
	y int
}

type Fold struct {
	dir   string
	index int
}

type FoldRes struct {
	newGrid [][]string
	newLen  int
}

func main() {
	data := constructData()
	yLen := data.yLen
	xLen := data.xLen
	foldedGrid := data.grid

	for i, fold := range data.folds {
		if i == 1 {
			fmt.Println("Part 1: ", countDots(foldedGrid)) // part 1 result
		}
		if fold.dir == "y" {
			res := foldY(foldedGrid, fold.index, yLen, xLen)
			yLen = res.newLen
			foldedGrid = res.newGrid
		} else {
			res := foldX(foldedGrid, fold.index, xLen, yLen)
			xLen = res.newLen
			foldedGrid = res.newGrid
		}
	}

	for _, row := range foldedGrid {
		fmt.Println(row) // Part 2 result
	}
}

func countDots(grid [][]string) int {
	count := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "#" {
				count++
			}
		}
	}
	return count
}

func foldX(grid [][]string, index, xLen, yLen int) FoldRes {
	var newGrid [][]string

	for y := 0; y < yLen; y++ {
		var row []string
		for i := 0; i < index; i++ {
			row = append(row, grid[y][i])
		}
		newGrid = append(newGrid, row)
	}

	for y := 0; y < yLen; y++ {
		for x := index + 1; x < xLen; x++ {
			if grid[y][x] == "#" {
				newGrid[y][int(math.Abs(float64(x)-float64(xLen-1)))] = grid[y][x]
			}
		}
	}

	return FoldRes{newGrid: newGrid, newLen: len(newGrid[0])}
}

func foldY(grid [][]string, index, yLen, xLen int) FoldRes {
	var newGrid [][]string

	for y := 0; y < index; y++ {
		var row []string
		for x := 0; x < xLen; x++ {
			row = append(row, grid[y][x])
		}
		newGrid = append(newGrid, row)
	}

	for y := index + 1; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if grid[y][x] == "#" {
				newGrid[int(math.Abs(float64(y)-float64(yLen-1)))][x] = grid[y][x]
			}
		}
	}

	return FoldRes{newGrid: newGrid, newLen: len(newGrid)}
}

func constructData() Data {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data []Coords
	var folds []Fold

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) > 1 {
			coords := Coords{x: parseInt(line[0]), y: parseInt(line[1])}
			input_data = append(input_data, coords)
		} else {
			r, _ := regexp.Compile(`[^y|x|0-9]`)
			s := r.ReplaceAllString(scanner.Text(), "")
			fold := Fold{dir: string(s[0]), index: parseInt(string(s[1:]))}
			folds = append(folds, fold)
		}
	}

	yLen := 0
	xLen := 0
	for i := len(folds) - 1; i >= 0; i-- {
		if folds[i].dir == "y" {
			yLen = folds[i].index*2 + 1
		} else {
			xLen = folds[i].index*2 + 1
		}
	}
	var grid [][]string
	for y := 0; y < yLen; y++ {
		var row []string
		for x := 0; x < xLen; x++ {
			row = append(row, " ")
		}
		grid = append(grid, row)
	}

	for _, spot := range input_data {
		grid[spot.y][spot.x] = "#"
	}

	return Data{grid: grid, folds: folds, yLen: yLen, xLen: xLen}
}

func parseInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
