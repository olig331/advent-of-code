package main

import "fmt"

type Scores struct{
	p1 int
	p2 int
}

func main() {
	rolls := 0
	restartCount := 0
	scores := Scores{p1: 0, p2:0}
	boardPos := Scores{p1: 1, p2:2}
	turn := "p1"


	for {	
		if scores.p1 >= 1000{
			fmt.Println("p1", scores.p1, scores.p2, rolls)
			fmt.Println(scores.p2 * (rolls + (restartCount * 1000)))
			break
		}
		if scores.p2 >= 1000 {
			fmt.Println("p2", scores.p2, scores.p1, rolls)
			fmt.Println(scores.p1 * (rolls + (restartCount * 1000)))
			break
		}
		count := 0
		if rolls == 999 {
			count = (rolls + 1) + 1 + 2
			rolls = 2
			restartCount ++
		} else {
			count = (rolls + 1) + (rolls + 2) + (rolls + 3)	
			rolls += 3
		}


		if turn == "p1" {
		 	for i := 0; i < count; i++ {
				boardPos.p1 ++ 
				if boardPos.p1 == 11 {
					boardPos.p1 = 1
				}
			}

			scores.p1 += boardPos.p1
			turn = "p2"
		}	else {
			for i := 0; i < count; i++ {
				boardPos.p2 ++ 
				if boardPos.p2 == 11 {
					boardPos.p2 = 1
				}
			}
			scores.p2 += boardPos.p2
			turn = "p1"
		}
	}
}
