/*
  looks for the highest scoring
  hourglass in a 6x6 2d array.
  xxx
   x   hourglass
  xxx
*/

package main

import (
	"fmt"
)

var highscore = -1000

func check(x, y int, a [6][6]int) int {
	var res int = 0
	var tmp int = y
	if x < 4 && y < 4 {
		for i := 0; i < 3; i++ {
			if i == 1 {
				res += a[tmp][x+1]
			} else {
				res += a[tmp][x] + a[tmp][x+1] + a[tmp][x+2]
			}
		  tmp += 1
		}
		return res
	}
	return -100
}

func main() {
	var arr = [6][6]int{}

	for i:= 0; i < len(arr); i++ {
		for j := 0; j < len(arr[j]); j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	var s int
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[y]); x++ {
			s = check(x,y,arr)
			if s > highscore {
				highscore = s
			}
		}
	}
	fmt.Println(highscore)
}
