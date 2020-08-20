package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aa() {
	fmt.Println("New Go Chapters")

}

func main() {

	rand.Seed(time.Now().UnixNano())

	aa()

}
