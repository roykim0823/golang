package main

import (
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int) // unbuffered channel
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
