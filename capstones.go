package main

import (
	"fmt"
	"math/rand"
	"time"
)

func capstone_2() {
	fmt.Println("New Go Chapters")
	fmt.Println("do page 103")
	fmt.Println(`
	don't avoid thinking when the opportunity 
	presents itself, few such opportunities of focused thinking 
	exist in life.
	`)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	capstone_2()

}
