package main

import "fmt"


func main(){
	data := constructData()
	fmt.Println(countLantonFish(80, sortLantonFish(data))) // Part 1
	fmt.Println(countLantonFish(256, sortLantonFish(data))) // Part 2
}


func sortLantonFish(data []int)[]int {
	sorted := []int{0,0,0,0,0,0,0,0,0}
	for _, val := range data {
		sorted[val] += 1
	}
	return sorted
}

func countLantonFish(days int, sortedLantons []int)int {
	copy := days;
	for {
		if(copy == 0){
			break
		}
		nextFish := sortedLantons[0]
		sortedLantons = append(sortedLantons, nextFish)
		sortedLantons = delete(sortedLantons, 0)		
		sortedLantons[6] += nextFish
		copy--
	}
	return count(sortedLantons)
}

func count(finalLantonCount []int)int {
	count := 0
	for _, val := range finalLantonCount {
		count += val
	} 
	return count
}


func constructData()[]int{
	data := []int{1,1,3,5,3,1,1,4,1,1,5,2,4,3,1,1,3,1,1,5,5,1,3,2,5,4,1,1,5,1,4,2,1,4,2,1,4,4,1,5,1,4,4,1,1,5,1,5,1,5,1,1,1,5,1,2,5,1,1,3,2,2,2,1,4,1,1,2,4,1,3,1,2,1,3,5,2,3,5,1,1,4,3,3,5,1,5,3,1,2,3,4,1,1,5,4,1,3,4,4,1,2,4,4,1,1,3,5,3,1,2,2,5,1,4,1,3,3,3,3,1,1,2,1,5,3,4,5,1,5,2,5,3,2,1,4,2,1,1,1,4,1,2,1,2,2,4,5,5,5,4,1,4,1,4,2,3,2,3,1,1,2,3,1,1,1,5,2,2,5,3,1,4,1,2,1,1,5,3,1,4,5,1,4,2,1,1,5,1,5,4,1,5,5,2,3,1,3,5,1,1,1,1,3,1,1,4,1,5,2,1,1,3,5,1,1,4,2,1,2,5,2,5,1,1,1,2,3,5,5,1,4,3,2,2,3,2,1,1,4,1,3,5,2,3,1,1,5,1,3,5,1,1,5,5,3,1,3,3,1,2,3,1,5,1,3,2,1,3,1,1,2,3,5,3,5,5,4,3,1,5,1,1,2,3,2,2,1,1,2,1,4,1,2,3,3,3,1,3,5}
	// data := []int{3,4,3,1,2}
	return data
}

func delete(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}