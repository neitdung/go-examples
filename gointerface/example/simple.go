package example

import (
	"fmt"
)

type Speaker interface {
    Speak() string
}

type Foo struct{}

func (Foo) Speak() string {
	return "Hello, I am Foo"
}

func SimpleRun() {
	var someSpeaker Speaker = Foo{}
	fmt.Println(someSpeaker.Speak())
}