package main

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func main(){

}

func TestArray01(t *testing.T){
	var x [100]int
	fmt.Printf("%d \n",x)
	x[99] = 42
	fmt.Printf("test %d \n",x[99])
	assert.Equal(t, 42,x[99])
}

