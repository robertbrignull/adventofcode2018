package day3

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type claim struct {
	id int
	rect rect
}

type rect struct {
	x int
	y int
	w int
	h int
}

func getInput() []claim {
	input := util.ReadFile("./day3/day3_input")
	input = strings.Replace(input, "\r\n", "\n", -1)
	lines := strings.Split(input, "\n")
	var claims []claim
	claimRegex := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	for _, line := range lines {
		match := claimRegex.FindStringSubmatch(line)
		if match == nil {
			log.Fatal("Unable to parse \"" + line + "\"")
		}
		id, _ := strconv.Atoi(match[1])
		x, _ := strconv.Atoi(match[2])
		y, _ := strconv.Atoi(match[3])
		w, _ := strconv.Atoi(match[4])
		h, _ := strconv.Atoi(match[5])
		rect := rect { x, y, w, h }
		claims = append(claims, claim { id, rect })
	}
	return claims
}

func getSize(claims []claim) (int, int) {
	w := 0
	h := 0
	for _, claim := range claims {
		if claim.rect.x + claim.rect.w + 1 > w { w = claim.rect.x + claim.rect.w + 1 }
		if claim.rect.y + claim.rect.h + 1 > h { h = claim.rect.y + claim.rect.h + 1 }
	}
	return w, h
}

func part1() {
	claims := getInput()

	w, h := getSize(claims)
	var grid = make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
	}

	for _, claim := range claims {
		for y := claim.rect.y; y < claim.rect.y+claim.rect.h; y++ {
			for x := claim.rect.x; x < claim.rect.x+claim.rect.w; x++ {
				grid[y][x] = grid[y][x] + 1
			}
		}
	}

	numOverlapping := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] > 1{
				numOverlapping += 1
			}
		}
	}
	fmt.Printf("part 1 result = %d\n", numOverlapping)
}

func overlaps(claim claim, allClaims []claim) bool {
	for _, other := range allClaims {
		if claim.id != other.id &&
		    claim.rect.x + claim.rect.w >= other.rect.x &&
			claim.rect.x < other.rect.x + other.rect.w &&
			claim.rect.y + claim.rect.h >+ other.rect.y &&
			claim.rect.y < other.rect.y + other.rect.h {
			return true
		}
	}
	return false
}

func part2() {
	claims := getInput()
	for _, claim := range claims {
		if !overlaps(claim, claims) {
			fmt.Printf("part 2 result = %d\n", claim.id)
			return
		}
	}
	log.Fatal("Unable to complete part 2")
}

func Run() {
	part1()
	part2()
}
