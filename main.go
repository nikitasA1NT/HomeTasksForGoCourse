package main

import (
	"HomeTasksForGoCourse/solutions/findmissingelem"
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
}
