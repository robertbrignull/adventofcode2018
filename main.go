package main

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/day1"
	"github.com/robertbrignull/adventofcode2018/day2"
	"github.com/robertbrignull/adventofcode2018/day3"
	"github.com/robertbrignull/adventofcode2018/day4"
	"log"
	"os"
	"strconv"
)

func runDay(dayNum int) {
	switch dayNum {
	case 1: day1.Run()
	case 2: day2.Run()
	case 3: day3.Run()
	case 4: day4.Run()
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		for i := 1; i <= 25; i++ {
			fmt.Printf("Day %d results...\n", i)
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
