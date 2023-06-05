package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(bubbleSort(x))
	fmt.Println(mergeSort(x))

	doubleArray := make([][]int, 5)

	for i := 0; i < 5; i++ {
		doubleArray[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			doubleArray[i][j] = rand.Intn(100)
		}
	}

	fmt.Println(doubleArray)
	fmt.Println(findMax(doubleArray))

	book := make(map[string]int)
	s := "MDCCCXXXIV"

	book["M"] = 1000
	book["D"] = 500
	book["C"] = 100
	book["X"] = 10
	book["V"] = 5
	book["I"] = 1

	fmt.Println(convertRomanToArabic(s, book))

}

func bubbleSort(arr []int) []int {
	sorted := true
	for i, _ := range arr {
		if i == 0 {
			continue
		}
		if arr[i-1] > arr[i] {
			sorted = false
			local := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = local
		}
	}

	if sorted {
		return arr
	} else {
		return bubbleSort(arr)
	}
}

func merge(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	arr := make([]int, (len(a) + len(b)))
	count := 0
	i := 0
	j := 0

	for count < (len(a) + len(b)) {

		if i == len(a) {
			arr[count] = b[j]
			j++
			count++
			continue
		}
		if j == len(b) {
			arr[count] = a[i]
			i++
			count++
			continue
		}

		if a[i] < b[j] {
			arr[count] = a[i]
			count++
			i++

		} else {
			arr[count] = b[j]
			count++
			j++
		}
	}

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	half := int(math.Floor(float64(len(arr) / 2)))

	sl1 := arr[:half]
	sl2 := arr[half:]

	return merge(mergeSort(sl1), mergeSort(sl2))
}

func findMax(arr [][]int) []int {
	max := 0
	index := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if arr[i][j] > max {
				index = i
				max = arr[i][j]
			}
		}
	}

	return arr[index]
}

func convertRomanToArabic(s string, book map[string]int) int {
	count := 0
	arr := strings.Split(s, "")

	for i := 0; i < len(arr); i++ {

		if arr[i] == "I" && i < (len(arr)-1) {
			if arr[i+1] != "I" {
				num := book[arr[i+1]] - book[arr[i]]
				count += num
				i++
				continue
			}
		}

		num, ok := book[arr[i]]
		if ok {
			count += num
		}
	}
	return count
}
