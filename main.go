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

	// Тесты Проверки последовательности
	PrintTestsOfSequenceCheck()

	// Тесты Циклической ротации
	PrintTestsOfCyclicRotation()
}

func PrintTestsOfFindMissingElem() {
	slice1 := []int{2, 3, 1, 5}
	fmt.Println("Source array:", slice1)
	result1, err1 := findmissingelem.Solution(slice1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Missing element:", result1)
	}

	slice2 := []int{7, 2, 3, 1, 5, 4}
	fmt.Println("Source array:", slice2)
	result2, err2 := findmissingelem.Solution(slice2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Missing element:", result2)
	}

	slice3 := []int{2, 3, 4, 5, 6}
	fmt.Println("Source array:", slice3)
	result3, err3 := findmissingelem.Solution(slice3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Missing element:", result3)
	}

	slice4 := []int{1, 2, 3, 4, 6}
	fmt.Println("Source array:", slice4)
	result4, err4 := findmissingelem.Solution(slice4)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println("Missing element:", result4)
	}

	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println("Source array:", slice5)
	result5, err5 := findmissingelem.Solution(slice5)
	if err5 != nil {
		fmt.Println("Error:", err5)
	} else {
		fmt.Println("Missing element:", result5)
	}

	slice6 := []int{}
	fmt.Println("Source array:", slice6)
	result6, err6 := findmissingelem.Solution(slice6)
	if err6 != nil {
		fmt.Println("Error:", err6)
	} else {
		fmt.Println("Missing element:", result6)
	}

	slice7 := []int{1}
	fmt.Println("Source array:", slice7)
	result7, err7 := findmissingelem.Solution(slice7)
	if err7 != nil {
		fmt.Println("Error:", err7)
	} else {
		fmt.Println("Missing element:", result7)
	}

	slice8 := []int{1, 2, 3, 5, 6, 7, 9}
	fmt.Println("Source array:", slice8)
	result8, err8 := findmissingelem.Solution(slice8)
	if err8 != nil {
		fmt.Println("Error:", err8)
	} else {
		fmt.Println("Missing element:", result8)
	}
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

func PrintTestsOfSequenceCheck() {
	slice1 := []int{4, 1, 3, 2}
	fmt.Println("Source array:", slice1)
	result1, err1 := sequencecheck.Solution(slice1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		if result1 == 1 {
			fmt.Println("Array is a sequence")
		} else {
			fmt.Println("Array is not a sequence")
		}
	}

	slice2 := []int{4, 1, 3}
	fmt.Println("Source array:", slice2)
	result2, err2 := sequencecheck.Solution(slice2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		if result2 == 1 {
			fmt.Println("Array is a sequence")
		} else {
			fmt.Println("Array is not a sequence")
		}
	}

	slice3 := []int{}
	fmt.Println("Source array:", slice3)
	result3, err3 := sequencecheck.Solution(slice3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		if result3 == 1 {
			fmt.Println("Array is a sequence")
		} else {
			fmt.Println("Array is not a sequence")
		}
	}

	slice4 := []int{1}
	fmt.Println("Source array:", slice4)
	result4, err4 := sequencecheck.Solution(slice4)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		if result4 == 1 {
			fmt.Println("Array is a sequence")
		} else {
			fmt.Println("Array is not a sequence")
		}
	}

	slice5 := []int{2}
	fmt.Println("Source array:", slice5)
	result5, err5 := sequencecheck.Solution(slice5)
	if err5 != nil {
		fmt.Println("Error:", err5)
	} else {
		if result5 == 1 {
			fmt.Println("Array is a sequence")
		} else {
			fmt.Println("Array is not a sequence")
		}
	}
}

func PrintTestsOfCyclicRotation() {
	slice1 := []int{3, 8, 9, 7, 6}
	k1 := 3
	fmt.Println("Source array:", slice1)
	result1 := cyclicrotation.Solution(slice1, k1)
	fmt.Println("Array right shifted", k1, "times:", result1)

	slice2 := []int{1, 2, 3, 4}
	k2 := 4
	fmt.Println("Source array:", slice2)
	result2 := cyclicrotation.Solution(slice2, k2)
	fmt.Println("Array right shifted", k1, "times:", result2)

	slice3 := []int{}
	k3 := 1
	fmt.Println("Source array:", slice3)
	result3 := cyclicrotation.Solution(slice3, k3)
	fmt.Println("Array right shifted", k3, "times:", result3)

	slice4 := []int{1}
	k4 := 2
	fmt.Println("Source array:", slice4)
	result4 := cyclicrotation.Solution(slice4, k4)
	fmt.Println("Array right shifted", k4, "times:", result4)

	slice5 := []int{1, 2}
	k5 := 0
	fmt.Println("Source array:", slice5)
	result5 := cyclicrotation.Solution(slice5, k5)
	fmt.Println("Array right shifted", k5, "times:", result5)
}
