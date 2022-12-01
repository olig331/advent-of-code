package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type WinnerReturnType struct {
	indexOfBoard int
	lastNum int
}

func main() {
	boards, _ := constructBoards()
	numbers := []int{67,31,58,8,79,18,19,45,38,13,40,62,85,10,21,96,56,55,4,36,76,42,32,34,39,89,6,12,24,57,93,47,41,52,83,61,5,37,28,15,86,23,69,92,70,27,25,53,44,80,65,22,99,43,66,26,11,72,2,98,14,82,87,20,73,46,35,7,1,84,95,74,81,63,78,94,16,60,29,97,91,30,17,54,68,90,71,88,77,9,64,50,0,49,48,75,3,59,51,33}

	// Part1	
	winningBoard := findWinner(boards, numbers)
	firstResult := calculateResultValue(boards[winningBoard.indexOfBoard], numbers[winningBoard.lastNum])
	fmt.Println(firstResult)

	// Part 2
	lastVal := 0
	var orderedWinList [][][]int
	for {
		if len(boards) == 0 {
			break;
		}
		nextWinner:= findWinner(boards, numbers)
		if nextWinner.indexOfBoard != -1 && nextWinner.lastNum != -1 {
			orderedWinList = append(orderedWinList, boards[nextWinner.indexOfBoard])
		}
		boards = removeBoard(boards, nextWinner.indexOfBoard)
		lastVal = nextWinner.lastNum
	}

	secondResult := calculateResultValue(orderedWinList[len(orderedWinList) - 1], numbers[lastVal])
	fmt.Println(secondResult)
}

func removeBoard(slice [][][]int, s int) [][][]int {
	return append(slice[:s], slice[s+1:]...)
}


func findWinner(boards [][][]int, numbers []int)(WinnerReturnType) {
	winner := -1
	lastValIndex := -1

	out:
	for i, num := range numbers {
		for j, board := range boards {
			for x, row := range board {
				if sum(row) == -5 {
					winner = j
					lastValIndex = i - 1		
					break out
				}
				for y, item := range row {
					col := []int{board[0][y], board[1][y], board[2][y], board[3][y], board[4][y]}
					if sum(col) == -5 {
						winner = j
						lastValIndex = i - 1
						break out
					} else {
						if num == item {
							boards[j][x][y] = -1
						}
					}
				}
			}
		}
	}	
	return WinnerReturnType{indexOfBoard: winner, lastNum: lastValIndex}
}

func calculateResultValue(board [][]int, lastNum int)int {
	count := 0
	for _, row := range board{
		for _, val := range row {
			if val != -1 {
				count += val
			} 			
		}
	}
	fmt.Println(count, lastNum)
	return count * lastNum
}


func constructBoards()([][][]int, error) {
	var boards [][][]int
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data [][]int
	for scanner.Scan() {

		line := scanner.Text()		
		arr := strings.Split(line, " ")
		var intArr []int
		for _, item := range arr {
			conv, _ := strconv.Atoi(item)
			intArr = append(intArr, conv)
		}
		input_data = append(input_data, intArr)
	}
	file.Close()
	var singleBoard [][]int
	for _, arr := range input_data {
		if len(arr) > 1 {
			singleBoard = append(singleBoard, arr)	
		} else {
			boards = append(boards, singleBoard)
			singleBoard = nil
		}
	}
	return boards, nil
}


func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
 }  