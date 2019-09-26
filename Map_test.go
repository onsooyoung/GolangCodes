package main

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func main(){

}

func Test_map(t *testing.T){

	var myMap map[int]string

	myMap = make(map[int]string)

	myMap[1] = "Harry"
	myMap[2] = "James"
	myMap[3] = "Dumbledore"
	assert.Equal(t, 3, len(myMap))

}

func Test_map_val(t *testing.T){

	var myMap map[int]string

	myMap = make(map[int]string)

	myMap[1] = "Harry"
	myMap[2] = "James"
	myMap[3] = "Dumbledore"
	assert.Equal(t, 3, len(myMap))

	a, b := myMap[1]

	assert.Equal(t, "Harry", a)
	assert.Equal(t, true, b)

	fmt.Printf("a=%s, b=%t \n", a, b)

}


