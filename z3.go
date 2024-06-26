package main

import (
	"fmt"
	"log"
)

func z3() {

	var c int
	fmt.Scan(&c)
	if c%2 == 0 {
		log.Println("четное")
	} else {
		log.Println("нечетное")
	}
}
