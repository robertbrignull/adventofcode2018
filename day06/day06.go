package day06

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
	infinite bool
}

type cell struct {
	nearest int
	distance int
}

type board struct {
	numPoints int
	points []point

	width int
	height int
	data [][]cell
}

func getInput() board {
	var points []point
	pointRegex := regexp.MustCompile(`^(\d+), (\d+)$`)
	for _, line := range strings.Split(util.ReadFile("./day06/day06_input"), "\n") {
		match := pointRegex.FindStringSubmatch(line)
		if match == nil {
			log.Fatal("Unable to parse \"" + line + "\"")
		}
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		points = append(points, point {
			x: x,
			y: y,
			infinite: false,
		})
	}

	numPoints := len(points)

	width := 0
	height := 0
	for _, point := range points {
		if point.x >= width {
			width = point.x + 1
		}
		if point.y >= height {
			height = point.y + 1
		}
	}

	data := make([][]cell, width)
	for x := 0; x < width; x++ {
		data[x] = make([]cell, height)
		for y := 0; y < height; y++ {
			data[x][y] = cell {
				nearest: -1,
				distance: -1,
			}
		}
	}

	for i, point := range points {
		data[point.x][point.y] = cell {
			nearest: i,
			distance: 0,
		}
	}

	return board { numPoints, points, width, height, data }
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func updatePoint(board *board, p, d, x, y int) bool {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return false
	}

	c := board.data[x][y]
	if c.distance == -1 {
		board.data[x][y] = cell {
			nearest: p,
			distance: d,
		}
		return true
	} else if c.distance == d && c.nearest != -1 {
		board.data[x][y] = cell {
			nearest: -1,
			distance: d,
		}
		return true
	}
	return false
}

func computeNearest(board *board) {
	maxD := max(board.width, board.height)
	for d := 1; d < maxD; d++ {
		madeChange := false
		for pIndex, p := range board.points {
			for dx := 0; dx < d; dx++ {
				if updatePoint(board, pIndex, d, p.x - d + dx, p.y - dx) { madeChange = true }
			}
			for dx := 0; dx < d; dx++ {
				if updatePoint(board, pIndex, d, p.x + dx, p.y - d + dx) { madeChange = true }
			}
			for dx := 0; dx < d; dx++ {
				if updatePoint(board, pIndex, d, p.x + d - dx, p.y + dx) { madeChange = true }
			}
			for dx := 0; dx < d; dx++ {
				if updatePoint(board, pIndex, d, p.x - dx, p.y + d - dx) { madeChange = true }
			}
		}
		if !madeChange {
			return
		}
	}
}

func findNearest(board *board, x, y int) int {
	nearest := -1
	bestD := -1
	for pIndex, p := range board.points {
		d := abs(x - p.x) + abs(y - p.y)
		if bestD == -1 || d < bestD {
			bestD = d
			nearest = pIndex
		}
	}
	return nearest
}

func computeInfinite(board *board) {
	d := max(board.width, board.height) * 2
	for pIndex, p := range board.points {
		if findNearest(board, p.x - d, p.y) == pIndex ||
			findNearest(board, p.x, p.y - d) == pIndex ||
			findNearest(board, p.x + d, p.y) == pIndex ||
			findNearest(board, p.x, p.y + d) == pIndex {
			board.points[pIndex].infinite = true
		}
	}
}

func part1() {
	board := getInput()
	computeNearest(&board)
	computeInfinite(&board)

	bestArea := 0
	for pIndex, p := range board.points {
		if p.infinite {
			continue
		}

		a := 0
		for x := 0; x < board.width; x++ {
			for y := 0; y < board.height; y++ {
				if board.data[x][y].nearest == pIndex {
					a += 1
				}
			}
		}
		if a > bestArea {
			bestArea = a
		}
	}
	fmt.Printf("part 1 result = %d\n", bestArea)
}

func part2() {
	board := getInput()

	numSafeLocations := 0
	for x := 0; x < board.width; x++ {
		for y := 0; y < board.height; y++ {
			d := 0
			for _, p := range board.points {
				d += abs(x - p.x)
				d += abs(y - p.y)
			}
			if d < 10000 {
				numSafeLocations += 1
			}
		}
	}
	fmt.Printf("part 2 result = %d\n", numSafeLocations)
}

func Run() {
	part1()
	part2()
}