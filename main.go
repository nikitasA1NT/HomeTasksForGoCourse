package main

import (
	"HomeTasksForGoCourse/solutions/findmissingelem"
	"HomeTasksForGoCourse/solutions/oddoccurrences"
	"fmt"
)

func main() {
	// Тесты Поиска отсутствующего элемента

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

	// Тесты Чудных вхождений в массив

	slice6 := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println("Source array:", slice6)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice6))

	slice7 := []int{9, 3, 9, 3, 9, 7}
	fmt.Println("Source array:", slice7)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice7))

	slice8 := []int{9, 3, 9, 3, 9, 9}
	fmt.Println("Source array:", slice8)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice8))

	slice9 := []int{1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Source array:", slice9)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice9))

	slice10 := []int{6, 1, 1, 2, 2, 3, 3}
	fmt.Println("Source array:", slice10)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice10))

	slice11 := []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 7, 7}
	fmt.Println("Source array:", slice11)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice11))

	slice12 := []int{}
	fmt.Println("Source array:", slice12)
	fmt.Println("Element without pair:", oddoccurrences.Solution(slice12))
}
