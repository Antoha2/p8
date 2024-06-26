package main

import "log"

type user struct {
	name string
	age  int
	city string
}

func z7() {

	c := &user{
		name: "Tom",
		age:  20,
		city: "London",
	}
	log.Println(c)
}
