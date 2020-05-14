package InterviewQuestionSolutions

import (
	"fmt"
	"math"
	"sort"

	"github.com/alazaralemayehu/datastructure"
)

func IsUniques(str string) {
	stringMap := make(map[string]bool)
	for _, v := range str {
		exists := stringMap[string(v)]
		if exists {
			fmt.Println("hey")
		}
		stringMap[string(v)] = true

	}
}

func ArePurmutation(string1, string2 string) bool {

	if len(string1) != len(string2) {
		return false
	}

	stringComparisonMap := make(map[rune]int)
	for _, s := range string1 {
		stringComparisonMap[s]++

	}

	for _, s := range string2 {
		_, exists := stringComparisonMap[s]
		if exists && stringComparisonMap[s] > 0 {
			stringComparisonMap[s]--
		} else {
			return false
		}

	}

	return true
}

func URLify(URLString string, lenghtOfString int) string {
	var stringRune []rune
	URLRune := []rune(URLString)
	for i := 0; i < lenghtOfString; i++ {
		if string(URLRune[i]) == (" ") {
			stringRune = append(stringRune, []rune("%20")...)
		} else {
			stringRune = append(stringRune, URLRune[i])
		}
	}
	return string(stringRune)
}

func PermutationPalindrom() {

}

func FindTheClosestValueINBST(bst *datastructure.BinarySearchTree, number int) int {
	return findTheClosetNode(bst.RootNode, number, math.MinInt64)
}
func findTheClosetNode(treeNode *datastructure.TreeNode, target int, closest int) int {
	if treeNode == nil {
		return closest
	}

	if math.Abs(float64(closest-target)) > math.Abs(float64(target-treeNode.Key)) {
		closest = treeNode.Key
	}

	if target < treeNode.Key {
		return findTheClosetNode(treeNode.LeftNode, target, closest)
	} else if target > treeNode.Key {
		return findTheClosetNode(treeNode.RightNode, target, closest)
	} else {
		return closest
	}
}

func SmartFibonacci(number int) {
	results := []int{0, 1}
	fibonacciNumbers := []int{0, 1}

	for i := 0; i < number; i++ {
		fibonacciResult := results[0] + results[1]
		fibonacciNumbers = append(fibonacciNumbers, fibonacciResult)

		// fmt.Print(fibonacciResult)
		results[0] = results[1]
		results[1] = fibonacciResult
		// fmt.Println(results)
	}

	for _, i := range fibonacciNumbers {
		fmt.Print(i, " ")
	}
}

func FibonaciWithMemoize(number int, memoize map[int]int) int {
	value, exists := memoize[number]
	fmt.Println(memoize)
	if exists {
		return value
	} else {
		memoize[number] = FibonaciWithMemoize(number-1, memoize) + FibonaciWithMemoize(number-2, memoize)
		return memoize[number]
	}
}

func ThreeSums(numbers []int, target int) (trippleSum [][]int) {
	datastructure.QuickSort(numbers)
	fmt.Println(numbers)
	for index, i := range numbers {
		left := index + 1
		right := len(numbers) - 1
		for left < right {
			numberLeft := numbers[left]
			numberRight := numbers[right]
			if numberLeft+numberRight+i > target {
				right--
			} else if numberRight+numberLeft+i < target {
				left++
			} else {
				t := []int{numberLeft, numberRight, i}
				trippleSum = append(trippleSum, t)
				left++
				right--
			}
			fmt.Println(left, right)

		}
	}
	return
}

func SmallestDifference(arrayOne, arrayTwo []int) (smallestPair [2]int) {
	datastructure.QuickSort(arrayOne)
	datastructure.QuickSort(arrayTwo)

	indexOne := 0
	indexTwo := 0
	smallest := math.MaxInt16
	current := math.MaxInt16
	for indexOne < len(arrayOne) && indexTwo < len(arrayTwo) {
		firstNumber := arrayOne[indexOne]
		secondNumber := arrayTwo[indexTwo]
		current = int(math.Abs(float64(firstNumber) - (float64(secondNumber))))
		fmt.Println(current)

		if firstNumber < secondNumber {
			indexOne++
		} else if secondNumber < firstNumber {
			indexTwo++
		} else {
			smallestPair[0] = firstNumber
			smallestPair[1] = secondNumber
			return
		}

		if current <= smallest {
			smallest = current
			smallestPair[0] = firstNumber
			smallestPair[1] = secondNumber
		}
	}
	return

}

func ValidateBST(bst *datastructure.BinarySearchTree) {
	fmt.Println(validateHelper(bst.RootNode))
	Slice
}

func validateHelper(treeNode *datastructure.TreeNode) bool {
	if treeNode.LeftNode == nil && treeNode.RightNode == nil {
		return true
	} else if treeNode.LeftNode == nil && treeNode.RightNode.Value < treeNode.Value {
		return false
	} else if treeNode.RightNode == nil && treeNode.LeftNode.Value > treeNode.Value {
		return false
	} else {
		if treeNode.RightNode.Value < treeNode.Value || treeNode.LeftNode.Value > treeNode.Value {
			return false
		} else {
			return validateHelper(treeNode.LeftNode) && validateHelper(treeNode.RightNode)
		}
	}

	return true
}

func Solution(A []int) int {
	sort.Ints(A)
	smallestIndex := 0
	largestIndex := len(A) - 1
	if A[largestIndex] < 0 || len(A) < 2 {
		if A[0] == 1 {
			return 2
		}
		return 1
	}
	if len(A) >= 2 {
		if A[0] > 1 {
			return 1
		}
	}

	smallestInteger := A[largestIndex] + 1

	for smallestIndex < largestIndex {

		firstNumber := A[smallestIndex]
		secondNumber := A[largestIndex]

		if firstNumber >= 0 {
			if A[smallestIndex+1]-firstNumber > 1 {
				smallestInteger = firstNumber + 1
				break
			}
		}

		if secondNumber-A[largestIndex-1] > 1 {
			fmt.Println(A)
			smallestInteger = A[largestIndex-1] + 1
		}
		largestIndex--
		smallestIndex++

	}
	if smallestInteger < 1 {
		return 1
	}

	return smallestInteger
}
