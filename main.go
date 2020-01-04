package main

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/day01"
	"github.com/robertbrignull/adventofcode2018/day02"
	"github.com/robertbrignull/adventofcode2018/day03"
	"github.com/robertbrignull/adventofcode2018/day04"
	"github.com/robertbrignull/adventofcode2018/day05"
	"github.com/robertbrignull/adventofcode2018/day06"
	"github.com/robertbrignull/adventofcode2018/day07"
	"github.com/robertbrignull/adventofcode2018/day08"
	"github.com/robertbrignull/adventofcode2018/day09"
	"log"
	"os"
	"strconv"
)

func runDay(dayNum int) {
	switch dayNum {
	case 1: day01.Run()
	case 2: day02.Run()
	case 3: day03.Run()
	case 4: day04.Run()
	case 5: day05.Run()
	case 6: day06.Run()
	case 7: day07.Run()
	case 8: day08.Run()
	case 9: day09.Run()
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
