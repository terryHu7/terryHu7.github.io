package main

import "fmt"

func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

func main() {
	fmt.Println(Binary([]int{1, 3, 5, 7, 9}, 9, 0, 4))
	fmt.Println(int(3 / 2))
}
