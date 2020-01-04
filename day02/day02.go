package day02

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"strconv"
	"strings"
)

func getInput() []string {
	input := util.ReadFile("./day02/day02_input")
	input = strings.Replace(input, "\r\n", "\n", -1)
	return strings.Split(input, "\n")
}

func part1() {
	input := getInput()

	numDoubles := 0
	numTriples := 0
	for _, s := range input {
		counts := make(map[rune]int)
		for _, c := range s {
			counts[c] = counts[c] + 1
		}

		isDouble := false
		isTriple := false
		for _, c := range s {
			if counts[c] == 2 {
				isDouble = true
			}
			if counts[c] == 3 {
				isTriple = true
			}
		}

		if isDouble {
			numDoubles += 1
		}
		if isTriple {
			numTriples += 1
		}
	}

	fmt.Printf("part 1 result = %d\n", numDoubles * numTriples)
}

func part2() {
	input := getInput()

	parts := make(map[string]string)
	for _, s := range input {
		for i := 0; i < len(s); i++ {
			part := s[:i] + s[i+1:]
			key := strconv.Itoa(i) + "_" + part
			if _, ok := parts[key]; ok {
				fmt.Printf("part 2 result = %s\n", part)
				return
			}
			parts[key] = s
		}
	}
	log.Fatal("Unable to complete part 2")
}

func Run() {
	part1()
	part2()
}
