package main

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func main() {
	fmt.Printf("%s\n", "Start Sort")
}

func FinderMaxNumber(sample []int) int {
	t := len(sample)
	var i = 0
	for {
		if i >= t {
			break
		}
		for j := 0; j < (t - 1); j++ {
			if sample[j] > sample[j+1] {
				tmp := sample[j]
				sample[j] = sample[j+1]
				sample[j+1] = tmp
			}
		}
		i++
	}

	fmt.Printf("Max Number : %d\n", sample[t-1])

	return sample[t-1]
}

func BubbleSort(sample []int, n int) []int {

	for i := 0; i < n-1; i++ {
		for j := 0; j < (n-1)-i; j++ {
			if sample[j] > sample[j+1] {
				tmp := sample[j]
				sample[j] = sample[j+1]
				sample[j+1] = tmp
			}
		}
	}

	return sample
}

func BubbleSortReCursive(sample []int, last int) []int {
	if last > 0 {
		for j := 0; j < last; j++ {
			if sample[j] > sample[j+1] {
				tmp := sample[j]
				sample[j] = sample[j+1]
				sample[j+1] = tmp
			}
		}

		BubbleSortReCursive(sample, last-1)
	}

	return sample
}

func SelectionSort(sample []int, n int)[]int{
	for i := 0; i < n; i++ {
		for j := 0 + i; j < (n-1); j++ {
			if sample[i] > sample[j+1] {
				tmp := sample[i]
				sample[i] = sample[j+1]
				sample[j+1] = tmp
			}
		}
	}

	return sample

}

func TestFinderMaxNumber(t *testing.T) {
	var data = []int{12, 78, 50, 11, 10, 20, 50, 20, 77, 99}

	maxNum := FinderMaxNumber(data)

	//https://godoc.org/gotest.tools/assert
	assert.Equal(t, maxNum, 99)
}

func TestBubbleSort(t *testing.T) {
	var data = []int{12, 78, 50, 11, 10, 20, 50, 20, 77, 1211}
	//fmt.Print(data)

	sortedData := BubbleSort(data, len(data))

	fmt.Printf("TestBubbleSort %d \n", data)

	assert.Equal(t, sortedData[0], 10)
	assert.Equal(t, sortedData[len(sortedData)-1], 1211)
}

func TestBubbleSortReCursive(t *testing.T) {
	var data = []int{12, 78, 50, -1, 11, 10, 20, 50, 20, 77, 11000}
	//fmt.Print(data)

	sortedData := BubbleSortReCursive(data, len(data)-1)

	fmt.Printf("TestBubbleSortReCursive %d \n", data)

	assert.Equal(t, sortedData[0], -1)
	assert.Equal(t, sortedData[len(sortedData)-1], 11000)
}


func TestSelectionSort(t *testing.T){
	var data = []int{12, 78, 50, -1, 11, 10, 20, 50, 20, 77, 11000, 111,22,532}
	//fmt.Print(data)

	sortedData := SelectionSort(data, len(data))

	fmt.Printf("TestSelectionSort %d \n", data)

	assert.Equal(t, sortedData[0], -1)
	assert.Equal(t, sortedData[len(sortedData)-1], 11000)
}