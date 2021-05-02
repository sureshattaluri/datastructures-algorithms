package algorithms

import (
	"fmt"
	"math"
)

func BubbleSort(nums []int) []int {
	//99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0
	change := true
	for change {
		change = false
		i := 0
		for i < len(nums)-1 {
			temp := nums[i]
			if temp > nums[i+1] {
				nums[i] = nums[i+1]
				nums[i+1] = temp
				change = true
			}
			i++
		}
		fmt.Printf("\n sorted nums by bubble sort: %v", nums)
	}
	return nums
}

func SelectionSort(nums []int) []int {
	//99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0
	for i := 0; i < len(nums); i++ {
		leastNumIndex := i
		leastNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < leastNum {
				leastNumIndex = j
				leastNum = nums[j]
			}
		}
		nums[leastNumIndex] = nums[i]
		nums[i] = leastNum
		fmt.Printf("\n sorted nums by selection sort: %v", nums)
	}
	return nums
}

func InsertionSort(nums []int) []int {
	//99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0
	for i := 1; i < len(nums); i++ {
		newNum := nums[i]
		for j := i-1; j >= 0; j-- {
			if nums[j] > newNum {
				nums[j+1] = nums[j]
				nums[j] = newNum
			} else {
				break
			}
		}
		fmt.Printf("\n sorted nums by insertion sort: %v", nums)
	}
	return nums
}


func MergeSort(nums []int) []int {
	//99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0
	length := len(nums)
	if length == 1 {
		return nums
	}
	middle := math.Floor(float64(length / 2))
	var left = nums[:int(middle)]
	var right = nums[int(middle):]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int,right []int) []int {
	leftIndex := 0;
	rightIndex := 0;
	result := []int{}
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	if leftIndex == len(left) {
		result = append(result, right[rightIndex:]...)
	} else {
		result = append(result, left[leftIndex:]...)
	}
	fmt.Printf("\n merged result %v", result)
	return result
}

