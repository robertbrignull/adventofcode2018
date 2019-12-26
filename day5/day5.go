package day5

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
)

type unit struct {
	value byte
	prev *unit
	next *unit
}

type polymer struct {
	first *unit
	last *unit
	length int
}

func getInput() polymer {
	input := util.ReadFile("./day5/day5_input")
	result := polymer {
		first: nil,
		last: nil,
		length: 0,
	}
	for _, v := range input {
		u := unit {
			value: byte(v),
			prev: result.last,
			next: nil,
		}
		if result.last == nil {
			result.first = &u
		} else {
			result.last.next = &u
		}
		result.last = &u
		result.length += 1
	}
	return result
}

func deleteUnit(p *polymer, u *unit) {
	if u.prev == nil {
		p.first = u.next
	} else {
		u.prev.next = u.next
	}
	if u.next == nil {
		p.last = u.prev
	} else {
		u.next.prev = u.prev
	}
	p.length -= 1
}

func deletePair(p *polymer, u, v *unit) {
	if u.next != v || v.prev != u {
		log.Fatal("u and v are not next to each other")
	}

	if u.prev == nil {
		p.first = v.next
	} else {
		u.prev.next = v.next
	}
	if v.next == nil {
		p.last = u.prev
	} else {
		v.next.prev = u.prev
	}
	p.length -= 2
}

func unitsReact(x, y byte) bool {
	return x == y + 32 || y == x + 32
}

func reactPolymer(polymer *polymer) {
	for unit := polymer.first; unit != nil && unit.next != nil; {
		if unitsReact(unit.next.value, unit.value) {
			deletePair(polymer, unit, unit.next)
			if unit.prev != nil {
				unit = unit.prev
			} else {
				unit = unit.next.next
			}
		} else {
			unit = unit.next
		}
	}
}

func part1() {
	polymer := getInput()
	reactPolymer(&polymer)
	fmt.Printf("part 1 result = %d\n", polymer.length)
}

func part2() {
	bestLength := -1
	for toRemove := byte(65); toRemove <= 90; toRemove++ {
		polymer := getInput()
		for unit := polymer.first; unit != nil; unit = unit.next {
			if unit.value == toRemove || unit.value == toRemove + 32 {
				deleteUnit(&polymer, unit)
			}
		}
		reactPolymer(&polymer)
		if bestLength == -1 || polymer.length < bestLength {
			bestLength = polymer.length
		}
	}
	fmt.Printf("part 2 result = %d\n", bestLength)
}

func Run() {
	part1()
	part2()
}
