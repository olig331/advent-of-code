package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Coords struct {
	x int
	y int
}

type Node struct {
	x int
	y int
	f int
	g int 
	h int
	cost int
	prev Coords
	neighbours []Coords
}



func main() {
	data := constructData()
	part1Grid := makeNodeGrid(data)
	part2Grid := makeNodeGrid(growGrid(data, 5))

	start := time.Now()
	part1Path := algo(Coords{y:0, x:0}, Coords{y: len(data) - 1, x: len(data[0]) - 1}, part1Grid)
	part2Path := algo(Coords{y:0, x:0}, Coords{y: len(data) * 5 - 1, x: len(data[0]) * 5 - 1}, part2Grid)

	fmt.Println(countPathCost(part1Path, part1Grid)) // part 1
	fmt.Println(countPathCost(part2Path, part2Grid)) // part 2
	elapsed := time.Since(start)
	fmt.Println("Took: ", elapsed) 
}

func countPathCost(path []Coords, grid [][]Node)int{
	count := 0
	for i:= 0; i < len(path); i++ {
		count += grid[path[i].y][path[i].x].cost
	}
	return count
}

func makeNodeGrid(data [][]int)[][]Node {
	rLen := len(data)
	cLen := len(data[0])
	var grid [][]Node
	for y := 0; y < len(data); y++ {
		var row []Node
		for x := 0; x < len(data[y]); x++ {
			item := Node{y:y, x:x, neighbours: getNeighbours(y, x, rLen - 1, cLen - 1), f: 0, g:100000000000000000, cost: data[y][x] }
			row = append(row, item)
		}
		grid = append(grid, row)		
	}

	return grid
}

// For Part 2
func growGrid(data [][]int, multiplyer int)[][]int {
	var newGrid [][]int
	rLen := len(data)
	cLen := len(data[0]) 

	for i := 0; i < rLen * multiplyer; i++ {
		var row []int
		for j := 0; j < cLen * multiplyer; j++ {
			val := data[i % len(data)][j % len(data[0])]
			xcycles := math.Floor(float64(j) / float64(cLen))
			ycycles := math.Floor(float64(i) / float64(rLen))
			total := xcycles + ycycles
			for x := 0; x < int(total); x++ {
				val ++
				if val == 10 {
					val = 1
				}
			}
			row = append(row, val) 
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

 
func algo(start Coords, end Coords, grid [][]Node)[]Coords {
	var openSet []Node
	grid[start.y][start.x].g = 0
	openSet = append(openSet, grid[start.y][start.x])
	lowestIndex := 0

	var cameFrom []Node
	cameFrom = append(cameFrom, grid[start.y][start.x])

	grid[start.y][start.x].g = 0 
	grid[start.y][start.x].f = h(grid[start.y][start.x], end, grid)		

	var finalPath []Coords
	
	for {
		if len(openSet) < 1 {
			break
		}

		for i := 0; i < len(openSet); i++ {
			if(i < len(openSet) - 1 ) {
				if openSet[i].f < openSet[lowestIndex].f {
					lowestIndex = i
				}
			}
		}

		curr := openSet[lowestIndex]

		openSet = remove(openSet, lowestIndex)

		if curr.x == end.x && curr.y == end.y {
			copy := Coords{y:curr.y, x:curr.x}
			for {
				if(copy.y == 0 && copy.x == 0) {
					break
				}
				finalPath = append(finalPath, copy)
				copy = grid[copy.y][copy.x].prev
			}
			return finalPath
		}
		
		for _, neighbour := range curr.neighbours {
			tempG := curr.g	+ grid[neighbour.y][neighbour.x].cost
			if tempG < grid[neighbour.y][neighbour.x].g {
				cameFrom = append(cameFrom, curr)
				grid[neighbour.y][neighbour.x].g = tempG
				grid[neighbour.y][neighbour.x].f = tempG + h(grid[neighbour.y][neighbour.x], end, grid)
				grid[neighbour.y][neighbour.x].prev = Coords{y:curr.y, x:curr.x}

				if !includes(openSet, grid[neighbour.y][neighbour.x]) {
					openSet = append(openSet, grid[neighbour.y][neighbour.x])
				}
			}
		}	
	}
	return finalPath
}

func remove(slice []Node, s int) []Node {
	return append(slice[:s], slice[s+1:]...)
}

func includes(visited []Node, curr Node)bool {
	for i := 0; i < len(visited); i++ {
		if reflect.DeepEqual(visited[i], curr){
			return true
		}
	}
	return false
}

func h(start Node, end Coords, grid[][]Node)int {
	return int(math.Abs(float64(start.x) - float64(end.x)) + math.Abs(float64(start.y) - float64(end.y)))
}


func getNeighbours(y, x, rLen, cLen int)[]Coords {
	var neighbours []Coords
	if y > 0 {neighbours = append(neighbours, Coords{y: y - 1, x:x})}
	if y < rLen {neighbours = append(neighbours, Coords{y: y + 1, x:x})}
	if x > 0 {neighbours = append(neighbours, Coords{y: y, x:x - 1})}
	if x < cLen {neighbours = append(neighbours, Coords{y: y, x:x + 1})}
	return neighbours
}

func constructData()[][]int {
	file, _ := os.Open("data.txt")	
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines)

	var input_data [][]int

	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		var row []int
		for _, char := range line {
			conv, _ := strconv.Atoi(char)
			row = append(row, conv)
		}
		input_data = append(input_data, row)
	}
	return input_data
}