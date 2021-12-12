package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int
var data = constructData()

type Coords struct {
	x int
	y int
}

var directions = []Coords{
	{x: -1, y: 1},
	{x: -1, y: -1},
	{x: 1, y: 1},
	{x: 1, y: -1},
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: -1, y: 0},
	{x: 1, y: 0},
}

func main() {
	cLen := len(data[0]) - 1
	rLen := len(data) - 1 
	
	for i := 0; i <= 100; i++ {
		
		for y := 0; y < len(data); y++{
			for x := 0; x < len(data[y]); x++ {
				data[y][x] += 1
			}
		}

		for y := 0; y < len(data); y++{
			for x := 0; x < len(data[y]); x++ {
				if data[y][x] >= 10 {
					flash(y, x, cLen, rLen)
				}
			}
		}

		for y := 0; y < len(data); y++{
			for x := 0; x < len(data[y]); x++ {
				if data[y][x] == -1 {
					data[y][x] = 0
				}
			}
		}
	}

	fmt.Println(count)
}


func flash(y, x, cLen, rLen int)  {
	data[y][x] = -1
	count += 1
	for _, newCoord := range directions {
		newX := x + newCoord.x
		newY := y + newCoord.y	
		if newY >= 0 && newY <= rLen && newX >= 0 && newX <= cLen  && data[newY][newX] != -1{
			data[newY][newX] += 1
			if data[newY][newX] >= 10 {
				flash(newX, newY, cLen, rLen)
			}
		}
	}
}



func constructData()[][]int{
	file, _ := os.Open("data.txt")	
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines)
	
	var input_data [][]int
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		var row []int
		for i := 0; i < len(line); i++ {
			conv, _ := strconv.Atoi(line[i])
			row = append(row, conv)
		}
		input_data = append(input_data, row)
	}
	return input_data
}