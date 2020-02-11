package myModule

import (
       "fmt"
)

func Version() {
       fmt.Println("Version 1.1.0")
}

func SayHello(name string) {
	fmt.Println("Hello " + name)
}

