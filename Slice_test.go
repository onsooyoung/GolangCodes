package main

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func main(){

}

func Test_slice(t *testing.T){

	var a []int
	a = []int{1,2,3,4}

	fmt.Printf("a=%d\n",a)

	assert.Equal(t, 4,len(a))

}

func Test_slice_make(t *testing.T){

	var b []int
	b = make([]int, 5, 7)

	assert.Equal(t, 5,len(b))
	assert.Equal(t, 7,cap(b))

	b = []int{1,2,3,4,5}

	fmt.Printf("b=%d, cap(b)=%d\n",b,cap(b))

	b = append(b, 6)

	fmt.Printf("b=%d, cap(b)=%d\n",b,cap(b))

	c := append(b, 7)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 8)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 9)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 10)

	assert.Equal(t, 10, cap(c))

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 11)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 12)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 13)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 14)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	c = append(c, 15)

	fmt.Printf("c=%d, cap(c)=%d\n",c,cap(c))

	assert.Equal(t, 20, cap(c))

}

func Test_slice_append(t *testing.T){

	var a []int
	a = []int{1,2,3,4}

	fmt.Printf("a=%d\n",a)

	a = append(a, 5)

	fmt.Printf("a len=%d, a cap=%d\n",len(a), cap(a))

	fmt.Printf("a=%d\n",a)

	assert.Equal(t, 5, len(a))

}