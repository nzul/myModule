package myModule

import (
	"fmt"
)

func Version() {
	fmt.Println("Version 1.2.0")
}

func SayHello(name string) {
	fmt.Println("Hello " + name)
}

func SayHi() {
	fmt.Println("Hi!")
}
