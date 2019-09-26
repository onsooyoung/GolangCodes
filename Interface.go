package main

import "fmt"

type Jam interface {
	GetOneSpoon() *SpoonofStrawberryJam
}


type Bread struct{
	val string
}

func (b *Bread) PutJam(jam *StrawberryJam){
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string{

	return "bread " +b.val
}


type StrawberryJam struct{
}


type SpoonofStrawberryJam struct{
}

func (s * SpoonofStrawberryJam) String() string  {

	return "+ strawberry"
}

func (j *StrawberryJam) GetOneSpoon() *SpoonofStrawberryJam{

	return &SpoonofStrawberryJam{}
}

func main(){

	bread := &Bread{}
	jam := &StrawberryJam{}

	bread.PutJam(jam)

	fmt.Printf("bread=%s",bread)
}


