package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main(){
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

	depth := 0
	horizontal := 0

	for _, line := range input_data {
		dirReg, _ := regexp.Compile("[a-z]+")
		valReg, _ := regexp.Compile("[0-9]+")
		val, _ := strconv.Atoi(valReg.FindString(line))
		dir := dirReg.FindString(line)

		if(dir == "up"){
			depth -= val
		}
		if(dir == "down"){
			depth += val
		}
		if(dir == "forward") {
			horizontal += val
		}
	}		
	fmt.Println(depth * horizontal)

	depth = 0
	horizontal = 0
	aim := 0

	for _, line := range input_data {
		dirReg, _ := regexp.Compile("[a-z]+")
		valReg, _ := regexp.Compile("[0-9]+")
		val, _ := strconv.Atoi(valReg.FindString(line))
		dir := dirReg.FindString(line)

		if(dir == "up"){
			aim -= val
		}
		if(dir == "down"){
			aim += val
		}
		if(dir == "forward") {
			horizontal += val
			depth += aim * val
		}
	}		

	fmt.Println(depth * horizontal)
}