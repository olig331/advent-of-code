package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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