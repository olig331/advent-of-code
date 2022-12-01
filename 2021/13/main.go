package main

/// INCOMPLETE

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	grid [][]int
	folds []string
	length Coords
}

type Coords struct {
	x int
	y int
}

func main() {
	data := constructData()
	grid := data.grid

	r, _ := regexp.Compile(`\w\=`)
	new := r.ReplaceAllString(data.folds[0], "")
	foldIndex, _ := strconv.Atoi(new)
	xL := len(grid[0]) - 1 

	// Fold on the x axis on the fold index
	for y := 0; y < len(grid); y++ {
		for x := foldIndex; x < len(grid[y]); x++ {
			grid[y][x] += grid[y][xL - x]
		}
	}

	count := 0
	for y := 0; y < len(grid); y++ {
		for x := foldIndex; x < len(grid[y]); x++ {
			if grid[y][x] > 0 {
				count++
			}
		}	
	}
	// for y := 0; y < len(grid); y++ {
	// 	for x := foldIndex + 1; x < len(grid[y]); x++ {
	// 		if grid[y][x] > 0 {
	// 			count++
	// 		}
	// 	}
	// }	
	fmt.Println("count %", count)
	// fmt.Println(foldDir, foldIndex)
	
}



func constructData()Data{
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	length := Coords{x:0, y:0}
	var folds []string
	var input_data []Coords
	for scanner.Scan(){
		r, _ := regexp.Compile(`^fold\salong\s`)
		matched := r.MatchString(scanner.Text())
		if matched {
			new := r.ReplaceAllString(scanner.Text(), "")
			folds = append(folds, new)
		} else {
			line := strings.Split(scanner.Text(), ",")
			xCoord, _ := strconv.Atoi(line[0])
			if xCoord > length.x {
				length.x = xCoord
			}
			yCoord, _ := strconv.Atoi(line[1])
			if yCoord > length.y {
				length.y = yCoord
			}
			coords := Coords{x:xCoord, y:yCoord} 
			input_data = append(input_data, coords)
		}
	}
	var grid [][]int
	for y := 0; y <= length.y; y++ {
		var row []int
		for x:= 0; x <= length.x; x++ {
			row = append(row, 0)			
		}
		grid = append(grid, row)
	}
	for _, item := range input_data {
		grid[item.y][item.x] = 1
	}

	return Data{ grid:grid, folds: folds, length: length }
}


// grid := [][]int{
// 	{1,0,1,1,0,0,1,0,0,1,0},
// 	{1,0,0,0,1,0,0,0,0,0,0},
// 	{0,0,0,0,0,0,1,0,0,0,1},
// 	{1,0,0,0,1,0,0,0,0,0,0},
// 	{0,1,0,1,0,0,1,0,1,1,1},
// 	{0,0,0,0,0,0,0,0,0,0,0},
// 	{0,0,0,0,0,0,0,0,0,0,0},
// }