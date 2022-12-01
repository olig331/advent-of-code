package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := constructData()
	sort.Ints(data)

	fmt.Println(calcLowestFuelCost(data, false)) // Part 1
	fmt.Println(calcLowestFuelCost(data, true)) // Part 2
}


func calcLowestFuelCost(data []int, scalingCost bool)int {
	lowest := data[0]
	highest := data[len(data) - 1]
	var positions []int
	for i := lowest; i <= highest; i++ {
		positions = append(positions, i)
	}
	count := 0
	var result int
	for i, position := range positions {
		count = 0
		for _, item := range data {
			fuelCost := position - item
			abs := int(math.Abs(float64(fuelCost)))		
			count += abs
			if scalingCost {
				for j := 1; j < abs; j ++ {
					count += j
				}
			}
		}
		if(i == 0){
			result = count
		} else if count < result {
			result = count
		}		
	}
	return result
}


func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
			if _, value := allKeys[item]; !value {
					allKeys[item] = true
					list = append(list, item)
			}
	}
	return list
}

func constructData()[]int {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, val := range line {
			conv, _ := strconv.Atoi(val)
			input_data = append(input_data, conv)
		}
	}
	file.Close()

	return input_data
}