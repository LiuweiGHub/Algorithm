package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 3, 2, 8, 1, 7, 4, 9, 0}
	arr = slectSort(arr)
	fmt.Println(arr)

}

//选择排序
func slectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		//关键是要假设最小数字的位置
		minPos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}

		//////////////////
		arr[minPos], arr[i] = arr[i], arr[minPos]
	}
	return arr
}
