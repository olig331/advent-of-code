package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
			return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

// INCOMPLETE

func main(){
	count := make(map[string]int)
	originalString := "NNCB"
	chars := strings.Split(originalString, "")
	inputs := constructData()
	pairs := []string{"NN", "NC", "CB"}
	pairCount := make(map[string]int)

	for _, val := range chars{
		count[val] += 1
	}

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < len(chars) - 1; j+=2 {
	// 		code := chars[j] + chars[j + 1]
	// 		chars = insert(chars, j + 1, inputs[code])
	// 		chars[j + 1] = inputs[code] 
	// 		count[inputs[code]] += 1
	// 	}
	// }

	// chars = strings.Split(originalString, "")	
	// fmt.Println(getResult(count))

	// for j := 0; j < len(chars) - 1; j+=2 { 
	// 	code := chars[j] + chars[j + 1]

	// }

	for i := 0; i < 40; i++ {
		for _, val := range pairs {
			split := strings.Split(val, "")	
			count[inputs[val]] += 1
			newPair1 := split[0] + inputs[val]
			newPair2 := split[1] + inputs[val]
			pairCount[newPair1] ++
			pairCount[newPair2] ++
		}
		fmt.Println(i)
	}
	
	fmt.Printf("%+v", count)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func getResult(count map[string]int)int {
	smallest := count["B"]
	largest := count["N"]
	for _, val := range count {
		if val < smallest {
			smallest = val
		}
		if val > largest {
			largest = val
		}
	}
	return largest - smallest
}

func constructData() map[string]string{

	file, _ := os.Open("test_data.txt")	
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines)
	
	input_data := make(map[string]string)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), " -> ")
		input_data[string(line[0])] = string(line[1])
	}

	return input_data
}