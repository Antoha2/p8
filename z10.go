package main

import "fmt"

func square(c chan int) {
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	num := <-c
	c <- num * num * num
}

func z10() {

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	squareChan <- testNum
	cubeChan <- testNum

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println(testNum, sum)

}
