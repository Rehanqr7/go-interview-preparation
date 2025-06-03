package main

import "fmt"

func main() {
	arr := []int{23, 54, 24, 1, 4, 3, 6, 90, 21, 87, 546, 42, 12, 45, 87, 1, 2, 7, 8, 0}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}
func merge(first []int, second []int) []int {
	mixed := make([]int, len(first)+len(second))
	i := 0
	j := 0
	k := 0
	for i < len(first) && j < len(second) {
		if first[i] > second[j] {
			mixed[k] = second[j]
			j++
		} else {
			mixed[k] = first[i]
			i++
		}
		k++
	}
	for i < len(first) {
		mixed[k] = first[i]
		k++
		i++
	}
	for j < len(second) {
		mixed[k] = second[j]
		k++
		j++
	}
	return mixed

}

func mergeSortWithIndex(arr []int, s, e int) {
	if e-s == 1 {
		return
	}
	mid := (s + e) / 2
	mergeSortWithIndex(arr, s, mid)
	mergeSortWithIndex(arr, mid, e)

	mergeInPlace(arr, mid, s, e)
}
func mergeInPlace(arr []int, mid, s, e int) {
	mix := make([]int, e-s)
	i := s
	j := mid
	k := 0
	for i < mid && j < e {
		if arr[i] > arr[j] {
			mix[k] = arr[j]
			j++
		} else {
			mix[k] = arr[i]
			i++
		}
		k++
	}
	for i < mid {
		mix[k] = arr[i]
		k++
		i++
	}
	for j < e {
		mix[k] = arr[j]
		k++
		j++
	}
	for l := 0; l < len(mix); l++ {
		arr[s+l] = mix[l]
	}

}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	s := low
	e := high
	mid := (low + high) / 2
	piviot := arr[mid]
	for s <= e {
		for arr[s] < piviot {
			s++
		}
		for arr[e] > piviot {
			e--
		}
		if s <= e {
			arr[e], arr[s] = arr[s], arr[e]
			s++
			e--
		}
	}
	quickSort(arr, low, e)
	quickSort(arr, s, high)

}
