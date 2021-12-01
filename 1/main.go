package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input_data []string
	for scanner.Scan() {
		input_data = append(input_data, scanner.Text())
	}
	file.Close()


	first_result := 0
	prev := 0
	for index, each_ln := range input_data {
		current, _ := strconv.Atoi(each_ln)
		if(index > 0) {
			if(current > prev){
				first_result ++
			}
		}
		prev = current
	}	

	fmt.Println(first_result)

	second_result := 0
	for index := range input_data {
		if(index >= 2 && index < len(input_data) - 1) {
			one, _ := strconv.Atoi(input_data[index - 2]) 
			two, _ := strconv.Atoi(input_data[index - 1]) 
			three, _ := strconv.Atoi(input_data[index])
			four, _ := strconv.Atoi(input_data[index + 1])

			first_set := one + two + three
			second_set := two + three + four

			if(second_set > first_set) {
				second_result ++
			}
		}
	
	}
	fmt.Println(second_result)
}