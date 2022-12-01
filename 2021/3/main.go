package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Count struct {
	zeros int
	ones int
}

type FirstQuestion struct {
	gamma string
	epsilon string
}

type FilterType struct {
	co2 []string
	oxygen []string
}

var testData = []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001", "00100", "01111", "00111", "00010", "01010"}

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

	binaryCount := countOccurances(input_data)

	firstResult := *&FirstQuestion{gamma: "", epsilon: ""}

	for _, count := range binaryCount {
		if(count.ones > count.zeros) {
			firstResult.gamma += "1"
			firstResult.epsilon += "0"
		} else {
			firstResult.gamma += "0"
			firstResult.epsilon += "1"
		}
	}
	gammaInt, _ := strconv.ParseInt(firstResult.gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(firstResult.epsilon,2 ,64)
	// Answer 1
	fmt.Println(gammaInt * epsilonInt)

	///////////////////////////////////////////////////////////////

	var oxygen []string
	var co2 []string

	for i := 0; i < 12; i++ {
		if i == 0 {		
			firstRun := filterOriginalList(input_data, 0, binaryCount)
			oxygen = firstRun.oxygen
			co2 = firstRun.co2
		} else {

			oxygenBinaryCount := countOccurances(oxygen)		

			oxyNum := "0"
			if oxygenBinaryCount[i].ones >= oxygenBinaryCount[i].zeros {
				oxyNum = "1"
			}

			co2DesiredCount := countOccurances(co2)
			co2Num := "1"
			if co2DesiredCount[i].ones >= co2DesiredCount[i].zeros {
				co2Num = "0"
			} 

			nextOxy := filterFilteredList(oxygen, i, oxyNum)
			nextCo2 := filterFilteredList(co2, i, co2Num)

			if len(nextOxy) > 0 {
				oxygen = nextOxy
			}

			if len(nextCo2) > 0 {
				co2 = nextCo2
			}
		}
	}
	fmt.Printf("%s\n%s\n", oxygen, co2)
	co2Rating, _ := strconv.ParseInt(string(co2[0]), 2, 64)
	oxygenRating, _ := strconv.ParseInt(string(oxygen[0]), 2, 64)

	fmt.Println(co2Rating * oxygenRating)
}


func filterFilteredList(list []string, resIndex int, desiredNum string) []string{
	var res []string
	for _, line := range list {
		if string(line[resIndex]) == desiredNum{
			res = append(res, line)
		}
	}
	return res
}


func filterOriginalList(input []string, resIndex int, prevResult []Count)FilterType{
	var co2 []string
	var oxygen []string

	for _, line := range input {
		if(prevResult[resIndex].ones >= prevResult[resIndex].zeros){
			if(string(line[resIndex]) == "1"){
				oxygen = append(oxygen, line)
			}else {
				co2 = append(co2, line)
			}
		} else {
			if(string(line[resIndex]) == "0"){
				oxygen = append(oxygen, line)
			} else {
				co2 = append(co2, line)
			}
		}
	}

	return FilterType{co2, oxygen}
}


func countOccurances(list []string) []Count{
	res := []Count{
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
		{zeros: 0, ones:0},
	}
	for _, line := range list {
		for i, char := range line {
			if(string(char) == "0"){
				res[i].zeros += 1
			} else {
				res[i].ones += 1
			}
		}
	}
	return res
}