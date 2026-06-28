package main

import "fmt"

func main() {

	var arr = [5]int{24, 34, 11, 56, 78}
	size := 4
	Val := 3751

	//right shift the elements of the array
	for i := size; i > 0; i-- {

		arr[i] = arr[i-1]

	}

	arr[0] = Val
	fmt.Println("Array after inserting an element at the beginning: ", arr)

}
