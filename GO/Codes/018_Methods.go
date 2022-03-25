package main

import (
	"fmt"
)

type Bird struct {
	kind string
	fly  bool
}

func (b *Bird) canFly() {
	if b.kind == "ostrich" {
		b.fly = false
	} else {
		b.fly = true
	}
}

func main() {
	someBird := Bird{kind: "ostrich"}
	someBird.canFly()
	if someBird.fly {
		fmt.Println("This bird can fly")
	} else {
		fmt.Println("This bird can't fly")
	}
}
