package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 3, 2, 8, 1, 7, 4, 9, 0}
	arr = slectSort(arr)
	fmt.Println(arr)

}

//ѡ�������㷨
func slectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		//�ؼ�
		minPos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}

		//////////////////



		////////////////////
		arr[minPos], arr[i] = arr[i], arr[minPos]
	}
	return arr
}

