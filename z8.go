package main

import "log"

func z8() {

	c := 0
	updateValue(&c)
	log.Println(c)
}

func updateValue(val *int) {
	*val = 1

}
