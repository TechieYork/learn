package main

import "fmt"

type Animal interface {
	Talk()
}

type Bird struct {
	Name string
	Age int
}

func NewBird(name string, age int) *Bird {
	return &Bird{
		Name: name,
		Age:  age,
	}
}

func (b *Bird) Talk() {
	fmt.Printf("Hi, I am a bird and my name is: %v\r\n", b.Name)
}

type Hawk struct {
	Bird				//This way to inherit from Bird
}

func NewHawk(name string, age int) *Hawk {
	return &Hawk{
		Bird{
			Name: name,
			Age:  age,
		},
	}
}

func (h *Hawk) Talk() {
	fmt.Printf("Hi, I am a hawk and my name is: %v\r\n", h.Name)
}


func main() {
	var a Animal
	a = NewBird("Polly", 5)
	a.Talk()

	b := NewBird("Dory", 6)
	b.Talk()

	h := NewHawk("William", 8)
	h.Talk()
}
