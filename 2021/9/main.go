package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main(){
	data := constructData()

	var lowest []int

	for i := 0; i < len(data); i++ {
		for j :=0; j < len(data[i]); j++ {
			var curr []int
			if i > 0 {
				curr = append(curr, data[i - 1][j])
			}
			if i < len(data) - 1 {
				curr = append(curr, data[i + 1][j])
			}

			if j > 0 {
				curr = append(curr, data[i][j - 1])
			}
			if j < len(data[i]) - 1 {
				curr = append(curr, data[i][j + 1])
			}
			sort.Ints(curr)
			if data[i][j] < curr[0] {
				lowest = append(lowest, data[i][j] + 1)
			}
		}
	}
	fmt.Println(count(lowest))

	for y := 0; y < len(data); y++{
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] < 9 {
				data[y][x] = 0
			}
		}
	}
	var results []int
	out:
	for y := 0; y < len(data); y++{
		for x := 0; x < len(data[y]); x++ {
			count := 0
			if(data[y][x] == 0) {
				count++
				res := checkNeighbours(data, y, x, count)
				data = res.data
				results = append(results, res.count)
				break out
			} 
		}
	}
	fmt.Printf("%+v", results)
}

type ReturnType struct {
	data [][]int
	count int
}

func checkNeighbours(data [][]int, y, x int, count int)ReturnType{
	copy := count
	for i := 0; i < len(directions); i++ {
		newX := x + directions[i].x
		newY := y + directions[i].y
		if newY >= 0 && newY < len(data) && newX >=0 && newX < len(data[newY]){
			if  data[y][x] == 0 {
				copy += 1
				data[y][x] = -1
				checkNeighbours(data, newY, newX, copy)
			}
		}
 	}
	return ReturnType{data:data, count:copy}
}


func count(arr []int)int {
	count := 0
	for i :=0; i < len(arr); i++ {
		count += arr[i] 
	}
	return count
}


func constructData()[][]int{
	file, _ := os.Open("test_data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data [][]int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		var row []int
		for _, val := range arr {
			conv, _ := strconv.Atoi(val)
			row = append(row, conv)
		}
		input_data = append(input_data, row)
	}
	file.Close()

	return input_data
}