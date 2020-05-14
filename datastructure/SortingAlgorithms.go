package datastructure

import "fmt"

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		k := 0
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
			k++
			fmt.Println(numbers)

		}
		fmt.Println()
		fmt.Println(numbers)

	}
}

func InsertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		k := 0
		for j := i; j > 0; j-- {
			if numbers[j] < numbers[j-1] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
			k++
			// fmt.Println(numbers)
		}
	}
	return numbers
}

func SelectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		var minimumnIndex int = i
		fmt.Println(numbers)
		for j := i; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				minimumnIndex = j
			}
		}
		if minimumnIndex != i {
			numbers[minimumnIndex], numbers[i] = numbers[i], numbers[minimumnIndex]
		}
	}
}

func MergeSort(numbers []int) {
	fmt.Println(Divide(numbers))
}

func Divide(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	var middle int = len(numbers) / 2
	leftSide := numbers[:middle]
	rightSide := numbers[middle:]

	newLeft := Divide(leftSide)
	newRight := Divide(rightSide)

	return Merge(newLeft, newRight)
}

func Merge(left, right []int) []int {
	size := len(left) + len(right)
	SortedNumbers := make([]int, size)

	for i, j, k := 0, 0, 0; k < size; k++ {
		if i == len(left) {
			SortedNumbers[k] = right[j]
			j++
		} else if j == len(right) {
			SortedNumbers[k] = left[i]
			i++
		} else if left[i] > right[j] {
			SortedNumbers[k] = right[j]
			j++
		} else if left[i] < right[j] {
			SortedNumbers[k] = left[i]
			i++
		} else {
			SortedNumbers[k] = left[i]
			i++
		}

	}

	return SortedNumbers
}
func QuickSort(elements []int) {
	quicSorter(elements, 0, len(elements)-1)
}

func quicSorter(elements []int, below int, upper int) {
	if below < upper {
		partition := DivideElements(elements, below, upper)
		quicSorter(elements, below, partition-1)
		quicSorter(elements, partition+1, upper)

	}
}

func DivideElements(elements []int, below int, upper int) int {
	pivotCounter := below
	center := elements[upper]
	for i := below; i < len(elements); i++ {
		if elements[i] < center {
			elements[pivotCounter], elements[i] = elements[i], elements[pivotCounter]
			pivotCounter++
		}
	}
	elements[pivotCounter], elements[upper] = elements[upper], elements[pivotCounter]
	return pivotCounter
}
