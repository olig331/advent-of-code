package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data := constructData()

	par1Vals := map[string]int{
		"]":57,
		")":3,
		"}":1197,
		">":25137,
	}

	part1CharCount := map[string]int{
		"]":0,
		"}":0,
		">":0,
		")":0,
	}	

	oppos := map[string]string {
		"(":")",
		"[":"]",
		"{":"}",
		"<":">",
		">":"<",
		"}":"{",
		"]":"[",
		")":"(",
	}
	open := []string{"<", "(", "{", "["}
	var part2Items [][]string
	var result []string
	for i := 0; i < len(data); i++ {
		var order []string
		out:
		for j := 0; j < len(data[i]); j++ {
			curr := data[i][j]
			if stringInSlice(curr, open) {
				order = append(order, curr)
			} else {
				if(curr == oppos[order[len(order) - 1]]) {
					order = removeItem(order, len(order) - 1)
				} else {
					fmt.Println()
					result = append(result, curr)
					break out
				}
			}
			if j == len(data[i]) -1 {
				part2Items = append(part2Items, data[i])
			}
		}
	}


	count := 0
	for _, val := range result {
		part1CharCount[val] ++
	}
	for key, val := range part1CharCount {
		count += (par1Vals[key] * val)
	}
	fmt.Printf("Part 1: %+v", count)


	// Part 2
	part2Vals := map[string]int{
		"]":2,
		")":1,
		">":4,
		"}":3,
	}
	var par2Result [][]string
	for _, row := range part2Items {
		var order []string
		for _, char := range row {
			if stringInSlice(char, open) {
				order = append(order, char)
			} else {
				if char == oppos[order[len(order) - 1]] {
					order = removeItem(order, len(order) -1)
				}
			}
		}
		var rowResult []string
		for x := len(order) -1; x >= 0; x-- {
			rowResult = append(rowResult, oppos[order[x]])
		}
		par2Result = append(par2Result, rowResult)
	}
	
	var intResult []int
	for _, row := range par2Result {
		count := 0
		for _, char := range row {
			count *= 5
			count += part2Vals[char]
		}
		intResult = append(intResult, count)
	}
	sort.Ints(intResult)
	fmt.Printf("Part 2 Reuslt: %+v", intResult[len(intResult) / 2])
}

func removeItem(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeRow(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
			if b == a {
					return true
			}
	}
	return false
}

func constructData() [][]string {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data [][]string
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		input_data = append(input_data, row)
	}
	file.Close()

	return input_data
}