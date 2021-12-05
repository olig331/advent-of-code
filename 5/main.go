package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Data struct {
	board [][]int
	input_data [][]int
	size int
}

func main(){
	startData, _ := constructData()
	board := startData.board
	size := startData.size
	input := startData.input_data
	board2, _ := constructData()

	fmt.Println(plotNumbers(board, input, size, false)) // Part 1
	fmt.Println(plotNumbers(board2.board, input, size, true)) // Part 2
}

func plotNumbers(board [][]int, input [][]int, size int, allowDiagonals bool)int{
	for i := 0; i < len(input); i++ {
		row := input[i]
		firstCoords := Coords{x: row[0], y: row[1]}
		secondCoords := Coords{x: row[2], y: row[3]}

		if(firstCoords.x == secondCoords.x) {
			board = calcLaterals(firstCoords.y, secondCoords.y, firstCoords.x, board, "y")
		}
		
		if(firstCoords.y == secondCoords.y) { 
			board = calcLaterals(firstCoords.x, secondCoords.x, firstCoords.y, board, "x")
		}

		if allowDiagonals {
			board = calcDiag(firstCoords, secondCoords, board)
		}
	}
	return count(board, size)
}

func calcLaterals(firstCoords int, secondCoords int, fixedPos int, board [][]int, axis string)[][]int {
	copy := firstCoords	
	for {
		if(copy == secondCoords){
			if(axis == "y"){ 
				board[copy][fixedPos] += 1; 
				break
			}
			board[fixedPos][copy] += 1
			break
		}
		if(axis == "y"){
			board[copy][fixedPos] += 1
		} else {
			board[fixedPos][copy] += 1
		}

		if(copy < secondCoords){ copy++ } 
		if( copy > secondCoords){ copy-- }		
	}		
	return board
}

func calcDiag(firstCoords Coords, secondCoords Coords, board[][]int)[][]int {
	if(firstCoords.y != secondCoords.y && firstCoords.x != secondCoords.x) { // Diagonals
		x := firstCoords.x
		y := firstCoords.y
		for {
			if(y == secondCoords.y || x == secondCoords.x){
				board[y][x] += 1
				break
			}
			board[y][x] += 1
			if(firstCoords.x < secondCoords.x && firstCoords.y > secondCoords.y){ x++; y-- }
			if(firstCoords.x > secondCoords.x && firstCoords.y > secondCoords.y){ x--; y-- }
			if(firstCoords.x > secondCoords.x && firstCoords.y < secondCoords.y){ x--; y++ }
			if(firstCoords.x < secondCoords.x && firstCoords.y < secondCoords.y){ x++; y++ }
		}		
	}
	return board
}

func count(board [][]int, size int)int {
	count := 0
	for  i := 0; i < size; i++{
    for j  := 0; j < size; j++ {
      if(board[i][j] > 1){
				count ++
      }
    }
  }
	return count
}

func constructData()(Data, error){
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	size := 0
	var input_data [][]int
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		var row []int
		for _, val := range arr {
			conv, _ := strconv.Atoi(val)
			if conv > size {
				size = conv
			}
			row = append(row, conv)
		}
		input_data = append(input_data, row)
	}
	file.Close()

	var board [][]int
	for  i := 0; i <= size; i++ {
		var row []int
		for j := 0; j <= size; j++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}
	return Data{board, input_data, size + 1}, nil
}