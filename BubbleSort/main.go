package main

import ("fmt"
	"math/rand"
)


var control bool


func bubbleSortIncrease(array [10]int) {
	for i:= 0; i < len(array); i++{
		control = true
		for j:=0; j < (len(array) - 1); j++{
			if array[j] > array[j+1] {
				array[j], array [j+1] = array[j+1], array[j]
				control = false
			}
		}
		if control {
			break
		}
	}
	for i:= 0; i < len(array); i++{
		fmt.Print(array[i], " ")
	}
}



func main() {

		var array [10]int
		for i:= 0; i < len(array); i++{
			array[i] = rand.Intn(1000)
		}
		fmt.Print("Generated numbers: ")
		for i := 0; i < len(array); i++{
			fmt.Print(array[i], " ")
		}

		fmt.Print("\nGrowing numbers: ")
		bubbleSortIncrease(array)







	//only for study :)
}