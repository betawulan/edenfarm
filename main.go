package main

import "fmt"

func maxValue(arr []int) int {
	var deretArray [][]int

	tempArray := []int{arr[0]}
	var tempIndex int
	for i := 1; i < len(arr)-1; i++ {
		if tempArray[tempIndex]+1 == arr[i] {
			tempArray = append(tempArray, arr[i])

			tempIndex++

			continue
		}
		if len(tempArray) > 1 {
			deretArray = append(deretArray, tempArray)

			tempArray = []int{arr[i]}
			tempIndex = 0
		}

	}

	var biggest int
	for _, value := range deretArray {
		// tempIndex := len(value)

		var count int
		var tempIndex int
		for i := 0; i < len(value); i++ {
			//if tempIndex > len(arr) {
			//	break
			//}

			if count == len(value)-1 {
				if value[len(value)-1] > biggest {
					biggest = value[len(value)-1]
				}
				count = 0
			}

			if arr[len(arr)-i-1] == value[tempIndex] {
				count++
				tempIndex++
			}

		}

	}

	return biggest
}
func main() {

	x := []int{1, 2, 3, 8, 9, 3, 2, 1}
	fmt.Println(maxValue(x))

}
