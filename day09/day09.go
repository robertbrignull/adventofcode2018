package day09

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"regexp"
	"strconv"
)

func getInput() (int, int) {
	inputRegex := regexp.MustCompile(`^(\d+) players; last marble is worth (\d+) points$`)
	input := util.ReadFile("./day09/day09_input")
	match := inputRegex.FindStringSubmatch(input)
	if match == nil {
		log.Fatal("Unable to parse \"" + input + "\"")
	}
	x, _ := strconv.Atoi(match[1])
	y, _ := strconv.Atoi(match[2])
	return x, y
}

type marble struct {
	value int
	next *marble
	prev *marble
}

func playGame(numPlayers int, finalMarbleScore int) int {
	scores := make([]int, numPlayers)
	currentPlayer := 0
	currentMarble := &marble { value: 0 }
	currentMarble.next = currentMarble
	currentMarble.prev = currentMarble
	for nextMarbleScore := 1; nextMarbleScore <= finalMarbleScore; nextMarbleScore++ {
		if nextMarbleScore % 23 != 0 {
			leftMarble := currentMarble.next
			rightMarble := leftMarble.next
			currentMarble = &marble {
				value: nextMarbleScore,
				prev: leftMarble,
				next: rightMarble,
			}
			leftMarble.next = currentMarble
			rightMarble.prev = currentMarble

		} else {
			scores[currentPlayer] += nextMarbleScore
			toRemove := currentMarble.prev.prev.prev.prev.prev.prev.prev
			scores[currentPlayer] += toRemove.value
			currentMarble = toRemove.next
			toRemove.prev.next = toRemove.next
			toRemove.next.prev = toRemove.prev
		}

		currentPlayer = (currentPlayer + 1) % numPlayers
	}

	winner := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[winner] {
			winner = i
		}
	}
	return scores[winner]
}

func Run() {
	numPlayers, finalMarbleScore := getInput()
	fmt.Printf("part 1 result = %d\n", playGame(numPlayers, finalMarbleScore))
	fmt.Printf("part 2 result = %d\n", playGame(numPlayers, finalMarbleScore * 100))
}
