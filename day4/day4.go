package day4

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	startsShift = iota
	fallsAsleep
	wakesUp
)

type event struct {
	timestamp int64
	guardNum int
	action int
}

func getInput() []event {
	layout := "2006-01-02 15:04"
	input := util.ReadFile("./day4/day4_input")
	input = strings.Replace(input, "\r\n", "\n", -1)
	lines := strings.Split(input, "\n")
	var result []event
	for _, line := range lines {
		timestampStr := line[1:17]
		timestamp, err := time.Parse(layout, timestampStr)
		if err != nil {
			log.Fatal("Unable to parse date \"" + timestampStr + "\"")
		}

		guardNum := -1
		action := -1
		if line[19:24] == "Guard" {
			guardNumStr := line[26:len(line)-13]
			guardNum, err = strconv.Atoi(guardNumStr)
			if err != nil {
				log.Fatal("Unable to parse guard num \"" + line[26:len(line)-13] + "\"")
			}
			action = startsShift
		} else if line[19:24] == "wakes" {
			action = wakesUp
		} else if line[19:24] == "falls" {
			action = fallsAsleep
		} else {
			log.Fatal("Unable to parse action for \"" + line + "\"")
		}

		result = append(result, event {
			timestamp: timestamp.Unix(),
			guardNum: guardNum,
			action: action,
		})
	}

	// Sort in to ascending order of timestamp
	sort.Slice(result, func(i, j int) bool {
		return result[i].timestamp < result[j].timestamp
	})

	// Fill in the missing guard numbers
	currentGuardNum := -1
	for i, event := range result {
		if event.action == startsShift {
			currentGuardNum = event.guardNum
		} else {
			result[i].guardNum = currentGuardNum
		}
	}

	// Verify no guard numbers are left unfilled
	for _, event := range result {
		if event.guardNum == -1 {
			log.Fatal("Could not populate guard numbers")
		}
	}

	return result
}

func countGuardSleeps(events []event) map[int]int64 {
	guardSleepCounts := make(map[int]int64)
	currentGuardNum := -1
	var currentSleepStart int64 = -1
	for _, event := range events {
		switch event.action {
		case startsShift:
			currentGuardNum = event.guardNum
			currentSleepStart = -1
		case fallsAsleep:
			currentSleepStart = event.timestamp
		case wakesUp:
			if currentSleepStart == -1 || currentGuardNum == -1 {
				log.Fatal("Woke up without falling asleep")
			}
			guardSleepCounts[currentGuardNum] = guardSleepCounts[currentGuardNum] + (event.timestamp - currentSleepStart)
		}
	}
	return guardSleepCounts
}

func findSleepiestGuard(events []event) int {
	guardSleepCounts := countGuardSleeps(events)

	sleepiestGuard := -1
	var mostSleep int64 = 0
	for guardNum, sleepCount := range guardSleepCounts {
		if sleepCount > mostSleep {
			sleepiestGuard = guardNum
			mostSleep = sleepCount
		}
	}
	return sleepiestGuard
}

func findSleepiestMinute(guardNum int, events []event) int {
	minutes := make([]int, 60)
	currentSleepStart := -1
	for _, event := range events {
		if event.guardNum != guardNum {
			continue
		}
		switch event.action {
		case fallsAsleep:
			currentSleepStart = int((((event.timestamp / 60) % 60) + 60) % 60)
			if currentSleepStart < 0 || currentSleepStart >= 60 {
				log.Fatal("sleep start \"" + strconv.Itoa(currentSleepStart) + "\" out of range")
			}
		case wakesUp:
			if currentSleepStart == -1 {
				log.Fatal("Woke up without falling asleep")
			}
			currentSleepEnd := int((((event.timestamp / 60) % 60) + 60) % 60)
			if currentSleepEnd < 0 || currentSleepEnd >= 60 {
				log.Fatal("sleep end \"" + strconv.Itoa(currentSleepEnd) + "\" out of range")
			}
			if currentSleepStart >= currentSleepEnd {
				log.Fatal("sleep end \"" + strconv.Itoa(currentSleepEnd) +
					"\" before sleep start \"" + strconv.Itoa(currentSleepStart) + "\"")
			}
			for minute := currentSleepStart; minute < currentSleepEnd; minute++ {
				minutes[minute] = minutes[minute] + 1
			}
		}
	}

	sleepiestMinute := -1
	mostAsleep := -1
	for minute, sleepiness := range minutes {
		if sleepiness > mostAsleep {
			sleepiestMinute = minute
			mostAsleep = sleepiness
		}
	}
	return sleepiestMinute
}

func part1() {
	events := getInput()
	sleepiestGuard := findSleepiestGuard(events)
	sleepiestMinute := findSleepiestMinute(sleepiestGuard, events)
	fmt.Printf("part 1 result = %d\n", sleepiestGuard * sleepiestMinute)
}

func Run() {
	part1()
}