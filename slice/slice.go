package main

import "fmt"

//slice 练习
func main() {
	// lenAndCap()
	// sliceForSlice(1, 2)
	// makeCreateSlice()
	appendSlice()
}

func lenAndCap() {
	//array
	var array = [10]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)
	//slice
	// var slice []int64
	slice := array[9:]
	fmt.Println(slice)
	var len = len(slice)
	var cap = cap(slice)
	fmt.Printf("len : %d  cap : %d \n", len, cap)
}

func sliceForSlice(a int64, b int64) string {
	var slice1 = []int64{1, 2, 3, 4, 5, 6, 7}
	var slice2 = slice1[0:5]

	fmt.Printf("slice1 \n  content:%v \n  len:%v \n  cap:%v \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 \n  content:%v \n  len:%v \n  cap:%v \n", slice2, len(slice2), cap(slice2))

	return "2"
}

func makeCreateSlice() {
	var slice = make([]int64, 10)
	slice[1] = 10
	slice = append(slice, 1)
	fmt.Printf("make create : \n  %v \n  len:%v \n  cap:%v \n", slice, len(slice), cap(slice))
}

func appendSlice() {
	var slice = make([]int64, 4)
	fmt.Printf("old slice : \n  %v \n  len:%v \n  cap:%v \n", slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("new slice : \n  %v \n  len:%v \n  cap:%v \n", slice, len(slice), cap(slice))

}
