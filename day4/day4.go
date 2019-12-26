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

	return result
}

type guard struct {
	num int
	totalSleep int
	sleepByMinute [60]int
}

func parseGuards(events []event) map[int]guard {
	guards := make(map[int]guard)
	currentGuardNum := -1
	currentSleepStart := -1
	for _, event := range events {
		switch event.action {
		case startsShift:
			currentGuardNum = event.guardNum
			currentSleepStart = -1
		case fallsAsleep:
			currentSleepStart = int((((event.timestamp / 60) % 60) + 60) % 60)
			if currentSleepStart < 0 || currentSleepStart >= 60 {
				log.Fatal("sleep start \"" + strconv.Itoa(currentSleepStart) + "\" out of range")
			}
		case wakesUp:
			if currentSleepStart == -1 || currentGuardNum == -1 {
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

			currentGuard, ok := guards[currentGuardNum]
			if !ok {
				var sleepByMinute [60]int
				currentGuard = guard {
					num: currentGuardNum,
					totalSleep: 0,
					sleepByMinute: sleepByMinute,
				}
			}
			currentGuard.totalSleep += currentSleepEnd - currentSleepStart
			for i := currentSleepStart; i < currentSleepEnd; i++ {
				currentGuard.sleepByMinute[i] += 1
			}

			guards[currentGuardNum] = currentGuard
		}
	}
	return guards
}

func part1(guards map[int]guard) {
	sleepiestGuardNum := -1
	sleepiestGuardValue := 0
	for _, guard := range guards {
		if guard.totalSleep > sleepiestGuardValue {
			sleepiestGuardNum = guard.num
			sleepiestGuardValue = guard.totalSleep
		}
	}

	sleepiestMinute := -1
	sleepiestMinuteValue := -1
	for minute, sleepiness := range guards[sleepiestGuardNum].sleepByMinute {
		if sleepiness > sleepiestMinuteValue {
			sleepiestMinute = minute
			sleepiestMinuteValue = sleepiness
		}
	}

	fmt.Printf("part 1 result = %d\n", sleepiestGuardNum * sleepiestMinute)
}

func Run() {
	events := getInput()
	guards := parseGuards(events)
	part1(guards)
}
