package main

import (
	"fmt"
	"math/rand"
	"time"
)

func functions_1() {
	fmt.Println("New Go Chapters")

	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	functions_1()

}
