package puppy

import (
	"fmt"
	"github.com/coderx31/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}

func From12() {
	fmt.Println("I'm from version 1.2.0")
}

func From13() {
	fmt.Println("I'm from version 1.3.0")
}

func From14() {
	fmt.Println("I'm from version 1.4.0")
}

func From15() {
	fmt.Println("I'm from version 1.5.0")
}
