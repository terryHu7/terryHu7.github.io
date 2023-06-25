package main

import (
	"fmt"
	"testing"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		var j int
		for j = i - 1; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func selectionSort2(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i; j < n; j++ { //j=i+1
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		maxIdx := i
		for j := 0; j <= i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	i, j := 1, len(arr)-1
	for i <= j {
		for i <= j && arr[i] < pivot {
			i++
		}
		for i <= j && arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	arr[0], arr[j] = arr[j], arr[0]
	quickSort2(arr[:j])
	quickSort2(arr[j+1:])
}

func quickSort3(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1

	for left < right {
		for arr[left] <= pivot && left < right {
			left++
		}

		for arr[right] >= pivot && left < right {
			right--
		}

		arr[left], arr[right] = arr[right], arr[left]
	}

	arr[0], arr[left] = arr[left], arr[0]

	quickSort3(arr[:left])
	quickSort3(arr[left+1:])
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left, pivot, right := partition(arr)

	res := append(append(quickSort(left), pivot), quickSort(right)...)
	return res
}

func partition(arr []int) (left []int, pivot int, right []int) {
	pivot = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	return left, pivot, right
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged
}

func TestSort(t *testing.T) {
	arr := []int{4, 1, 2, 3, 5, 7, 6, 10}

	quickSort3(arr)
	fmt.Println(arr)
	// bubbleSort(arr)

	// insertionSort(arr)

	// quickSort(arr)

	// mergeSort(arr)
}

func TestTmp(t *testing.T) {
	arr := []int{0, 1, 2}
	tmp := []int{}
	tmp = append(tmp, arr[:0]...)
	fmt.Println(tmp)
}
