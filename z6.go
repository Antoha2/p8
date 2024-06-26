package main

import "log"

func z6() {
	slice := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, d := range slice {
		sum = sum + d
	}
	log.Println(sum)
}
