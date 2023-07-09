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
