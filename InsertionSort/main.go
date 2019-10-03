package main

import (
	"fmt"
)


var aux,j int

func insertionSort(array [9]int) {

	//array = {8, 2, 4, 9, 3, 6, 1, 10, 7}
	for i := 1; i < len(array); i++ {
		aux = array[i]                  // aux = 2
		j = i - 1                       // j = 0
		for j >= 0 && array[j] > aux {  // (0 >= 0) e (8 > 2)
			array[j+1] = array[j]       // array[j+1] = 8
			j = j - 1                   // j = -1
		}
		array[j+1] = aux                // array[j+1] = 2;
	}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
}


func main() {

	array := [9]int{8, 2, 4, 9, 3, 6, 1, 10, 7}

	fmt.Print("Numbers: ")
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Print("\nNumbers in Order: ")
	insertionSort(array)










	//only for study :)
}