package main

import (
	"HomeTasksForGoCourse/solutions/findmissingelem"
	"HomeTasksForGoCourse/solutions/oddoccurrences"
	"fmt"
)

func main() {
	// Тесты Поиска отсутствующего элемента
	PrintTestsOfFindMissingElem()

	// Тесты Чудных вхождений в массив
	PrintTestsOfOddOccurrences()
}

func PrintTestsOfFindMissingElem() {
	slice1 := []int{2, 3, 1, 5}
	fmt.Println(findmissingelem.Solution(slice1))

	slice2 := []int{7, 2, 3, 1, 5, 4}
	fmt.Println(findmissingelem.Solution(slice2))

	slice3 := []int{2, 3, 4, 5, 6}
	fmt.Println(findmissingelem.Solution(slice3))

	slice4 := []int{1, 2, 3, 4, 6}
	fmt.Println(findmissingelem.Solution(slice4))

	slice5 := []int{}
	fmt.Println(findmissingelem.Solution(slice5))
}

func PrintTestsOfOddOccurrences() {
	slice6 := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println("Source array:", slice6)
	result6, err6 := oddoccurrences.Solution(slice6)
	if err6 != nil {
		fmt.Println("Error:", err6)
	} else {
		fmt.Println("Element without pair:", result6)
	}

	slice7 := []int{9, 3, 9, 3, 9, 7}
	fmt.Println("Source array:", slice7)
	result7, err7 := oddoccurrences.Solution(slice7)
	if err7 != nil {
		fmt.Println("Error:", err7)
	} else {
		fmt.Println("Element without pair:", result7)
	}

	slice8 := []int{9, 3, 9, 3, 9, 9}
	fmt.Println("Source array:", slice8)
	result8, err8 := oddoccurrences.Solution(slice7)
	if err8 != nil {
		fmt.Println("Error:", err8)
	} else {
		fmt.Println("Element without pair:", result8)
	}

	slice9 := []int{1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Source array:", slice9)
	result9, err9 := oddoccurrences.Solution(slice9)
	if err9 != nil {
		fmt.Println("Error:", err9)
	} else {
		fmt.Println("Element without pair:", result9)
	}

	slice10 := []int{6, 1, 1, 2, 2, 3, 3}
	fmt.Println("Source array:", slice10)
	result10, err10 := oddoccurrences.Solution(slice10)
	if err10 != nil {
		fmt.Println("Error:", err10)
	} else {
		fmt.Println("Element without pair:", result10)
	}

	slice11 := []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 7, 7}
	fmt.Println("Source array:", slice11)
	result11, err11 := oddoccurrences.Solution(slice11)
	if err11 != nil {
		fmt.Println("Error:", err11)
	} else {
		fmt.Println("Element without pair:", result11)
	}

	slice12 := []int{}
	fmt.Println("Source array:", slice12)
	result12, err12 := oddoccurrences.Solution(slice12)
	if err12 != nil {
		fmt.Println("Error:", err12)
	} else {
		fmt.Println("Element without pair:", result12)
	}

	slice13 := []int{1}
	fmt.Println("Source array:", slice13)
	result13, err13 := oddoccurrences.Solution(slice13)
	if err13 != nil {
		fmt.Println("Error:", err13)
	} else {
		fmt.Println("Element without pair:", result13)
	}
}
