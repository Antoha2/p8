package main

import "log"

func z4() {
	i := 1
	for {
		log.Println(i)
		if i == 5 {
			break
		}
		i++
	}
}
