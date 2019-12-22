package day1

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"strconv"
	"strings"
)

func getInput() []int {
	input := util.ReadFile("./day1/day1_input")
	var result []int
	for _, line := range strings.Split(input, "\n") {
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, x)
	}
	return result
}

func part1(input []int) {
	sum := 0
	for _, x := range input {
		sum += x
	}
	fmt.Printf("part 1 result = %d\n", sum)
}

func part2(input []int) {
	sum := 0
	sumsSeen := make(map[int]bool)
	for {
		for _, x := range input {
			_, ok := sumsSeen[sum]
			if ok {
				fmt.Printf("part 2 result = %d\n", sum)
				return
			}
			sumsSeen[sum] = true
			sum += x
		}
	}
}

func Run()  {
	input := getInput()
	part1(input)
	part2(input)
}
