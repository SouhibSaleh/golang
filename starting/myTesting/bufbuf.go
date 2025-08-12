package test

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int { return len(p) }
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func TypeOnly() {
	arr := make(IntSlice, 0)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 1)
	arr = append(arr, 2)
	sort.Sort(arr)
	fmt.Println(arr)
	sort.Sort(sort.Reverse(arr))
	fmt.Println(arr)
}

func BufbufPlayground() {

}
