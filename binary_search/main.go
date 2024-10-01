package main

import "fmt"

func main() {

	out := basic_binary_search([]int{1, 2, 3, 4, 5, 6, 7}, 6)
	fmt.Println("Output :: ", out)

	fmt.Println("order agnostics", order_agnostic_binary_search([]int{8, 7, 6, 5, 4, 3}, 5))
}

func basic_binary_search(arr []int, key int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)>>2
		switch {
		case arr[mid] == key:
			return mid
		case arr[mid] < key:
			left = mid + 1
		case arr[mid] > key:
			right = mid - 1
		}

	}
	return -1
}

func order_agnostic_binary_search(arr []int, key int) int {

	left := 0
	right := len(arr) - 1

	isAsc := arr[left] < arr[right]

	for left <= right {
		mid := left + (right-left)>>2

		if arr[mid] == key {
			return mid
		}
		if isAsc {
			if arr[mid] < key {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if arr[mid] < key {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return -1
}
