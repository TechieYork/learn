package main

import (
	"fmt"
	"sort"
)

func testInitArray() {
	fmt.Println("====== Test Init Array ======")

	//Init empty array
	var arrayEmpty [5]int
	fmt.Println(arrayEmpty)

	arrayEmpty1 := [5]int{}
	fmt.Println(arrayEmpty1)

	//Init array with value
	var arrayWithValue [5]int
	arrayWithValue = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayWithValue)

	arrayWithValue1 := [...]int{1, 2, 3, 4, 5}		//The size will automaticlly set
	fmt.Println(arrayWithValue1)

	arrayWithValue2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayWithValue2)

	//Init multi dimensional array
	var arrayOfTwoDimensionEmpty [2][3]int
	fmt.Println(arrayOfTwoDimensionEmpty)

	arrayOfTwoDimensionEmpty1 := [2][3]int{}
	fmt.Println(arrayOfTwoDimensionEmpty1)

	//Init multi dimensional array with value
	var arrayOfTwoDimensionWithValue [2][3]int
	arrayOfTwoDimensionWithValue = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arrayOfTwoDimensionWithValue)

	arrayOfTwoDimensionWithValue1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arrayOfTwoDimensionWithValue1)
}

func testInitSlice() {
	fmt.Println("====== Test Init Slice ======")

	//Init slice
	var sliceEmpty []int
	fmt.Println(sliceEmpty)

	sliceEmpty1 := []int{}
	fmt.Println(sliceEmpty1)

	sliceEmpty2 := make([]int, 0)
	fmt.Println(sliceEmpty2)

	//Init slice with value
	var sliceWithValue []int
	sliceWithValue = []int{1, 2 ,3 ,4, 5}
	fmt.Println(sliceWithValue)

	sliceWithValue1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceWithValue1)

	//Init slice with length and capacity
	sliceWithLenAndCap := make([]int, 10, 100)   //A 100 capacity slice with initial 10 length
	fmt.Println(sliceWithLenAndCap)
}

func testCopy() {
	fmt.Println("====== Test Copy ======")

	//Array and slice are complicated data structure
	//When copy, please note that, using 'copy' function provided by standard library to do a deep copy

	//Test copy without using 'copy'
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Println("Original slice1: ", slice1)
	fmt.Println("New slice from slice2 := slice1: ", slice2)
	slice1[0] = 99
	fmt.Println("After set slice1[0] = 99, Original slice1: ", slice1)
	fmt.Println("After set slice1[0] = 99, New slice2: ", slice2)

	//Test copy with using 'copy'
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{0, 0, 0, 0, 0}
	copy(slice4, slice3)
	fmt.Println("Original slice3: ", slice3)
	fmt.Println("New slice from copy(slice4, slice3): ", slice4)
	slice3[0] = 99
	fmt.Println("After set slice3[0] = 99, Original slice3: ", slice3)
	fmt.Println("After set slice3[0] = 99, New slice4: ", slice4)
}

func testLenAndCap() {
	fmt.Println("====== Test len and cap ======")

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v, Len:%v, Cap:%v\r\n", slice, len(slice), cap(slice))

	slice1 := make([]int, 5, 20)
	fmt.Printf("Slice: %v, Len:%v, Cap:%v\r\n", slice1, len(slice1), cap(slice1))
}

func testCreateSlice() {
	fmt.Println("====== Test slice ======")

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v, Len:%v, Cap:%v\r\n", array, len(array), cap(array))

	//Create slice
	slice := array[:4]			//Means from begin to 4 (4 not included)
	slice1 := array[1:4]		//Means from index 1 to 4 (4 nod included)
	slice2 := array[2:]			//Means from index 2 to end (The last one included)

	fmt.Println("After [:4]: ", slice)
	fmt.Println("After [1:4]: ", slice1)
	fmt.Println("After [2:]: ", slice2)

	//Modify slice will change the array as well
	slice[0] = 99
	fmt.Printf("Array: %v, Len:%v, Cap:%v\r\n", array, len(array), cap(array))
	slice[0] = 1

	//Modify slice without modifying the array
	slice3 := make([]int, 3)
	copy(slice3, array[1:4])
	slice3[0] = 99
	fmt.Println("After copy from[1:4]: ", slice3)
	fmt.Printf("After copy from[1:4] and set [0] = 99, Array: %v, Len:%v, Cap:%v\r\n", array, len(array), cap(array))
}

func testAppend() {
	fmt.Println("====== Test Append ======")

	slice := []int{}
	slice = append(slice, 1, 2, 3)
	fmt.Println("After append: ", slice)
}

func testSearch() {
	fmt.Println("====== Test Search ======")

	//Search using sort.Search
	slice := []int{1, 2, 3, 4, 5}
	x := 2
	index := sort.Search(len(slice), func (i int) bool {return slice[i] >= x})

	if index < len(slice) && slice[index] == 2 {
		//Using slice[index] == 2 to check whether 2 is found in the slice
		//Because sort.Search could only accept >= or <= operator
		fmt.Println("sort.Search: Found 2 in slice")
	} else {
		fmt.Println("sort.Search: Not found 2 in slice")
	}

	//Search using sort.SearchInt
	index = sort.SearchInts(slice, 2)

	if index < len(slice) {
		fmt.Println("sort.SearchInts: Found 2 in slice")
	} else {
		fmt.Println("sort.SearchInts: Not found 2 in slice")
	}
}

func testSort() {
	fmt.Println("====== Test Sort ======")

	//Sort using sort.SliceStable
	slice := []int{3, 4, 1, 5, 2}
	sort.SliceStable(slice, func (i, j int) bool {return slice[i] > slice[j]})
	fmt.Println("sort.SliceStable: ", slice)

	//Sort using sort.Sort
	sort.Sort(sort.IntSlice(slice))
	fmt.Println("sort.Sort: ", slice)

	//Sort using sort.Stable
	sort.Stable(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println("sort.Stable: ", slice)

	//Sort using sort.Ints
	sort.Ints(slice)
	fmt.Println("sort.Ints: ", slice)
}

func main() {
	testInitArray()
	testInitSlice()
	testCopy()
	testLenAndCap()
	testCreateSlice()
	testAppend()
	testSearch()
	testSort()
}
