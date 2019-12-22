package main

import (
	"github.com/robertbrignull/adventofcode2018/day1"
	"github.com/robertbrignull/adventofcode2018/day2"
	"log"
	"os"
	"strconv"
)

func runDay(dayNum int) {
	switch dayNum {
	case 1: day1.Run()
	case 2: day2.Run()
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		for i := 0; i < 25; i++ {
			runDay(i)
		}
	} else if len(args) == 2 {
		dayNum, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		runDay(dayNum)
	} else {
		log.Fatal("Usage: go run main.go [day-num]")
	}
}
