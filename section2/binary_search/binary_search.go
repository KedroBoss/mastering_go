// Package binarysearch does binary search
// Pseudocode:
// Let min = 0 and max = n-1.
// Compute guess as the average of max and min, rounded down (so that it is an integer).
// If array[guess] equals target, then stop. You found it! Return guess.
// If the guess was too low, that is, array[guess] < target, then set min = guess + 1.
// Otherwise, the guess was too high. Set max = guess - 1.
// Go back to step 2.
package main

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func binarySearch(array []int, target, minIndex, maxIndex int) int {
	if maxIndex < minIndex {
		return -1
	}
	mid := int((minIndex + maxIndex) / 2)
	switch {
	case array[mid] > target:
		return binarySearch(array, target, minIndex, mid)
	case array[mid] < target:
		return binarySearch(array, target, mid+1, maxIndex)
	default:
		return mid
	}
}

func iterBinarySearch(array []int, target, minIndex, maxIndex int) int {
	mini := minIndex
	maxi := maxIndex
	var mid int
	for mini < maxi {
		mid = int((mini + maxi) / 2)
		if array[mid] > target {
			maxi = mid
		} else if array[mid] < target {
			mini = mid
		} else {
			return mid
		}
		fmt.Printf("%v,%v,%v\n", mid, mini, maxi)
	}
	return -1
}

func main() {
	// DO NOT USE NON-EXISTING 
	a := binarySearch(primes, 7, 0, 10)
	fmt.Println(a)
}
