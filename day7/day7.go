package day7

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"log"
	"regexp"
	"strings"
)

func getInput() (map[int][]int, int) {
	result := make(map[int][]int)
	max := 0
	lineRegex := regexp.MustCompile(`^Step (.) must be finished before step (.) can begin.$`)
	for _, line := range strings.Split(util.ReadFile("./day7/day7_input"), "\n") {
		match := lineRegex.FindStringSubmatch(line)
		if match == nil {
			log.Fatal("Unable to parse \"" + line + "\"")
		}
		a := int(match[1][0] - 'A')
		b := int(match[2][0] - 'A')
		if a < 0 || a >= 26 || b < 0 || b >= 26 {
			log.Fatal("Found invalid letters in \"" + line + "\"")
		}
		result[b] = append(result[b], a)
		if a > max { max = a }
		if b > max { max = b }
	}
	return result, max + 1
}

type worker struct {
	currentStep int
	timeComplete int
}

// find the next worker that will become available
func findNextWorker(workers []worker, onlyExecutingWorkers bool) *worker {
	earliestWorker := -1
	for i, worker := range workers {
		if (!onlyExecutingWorkers || worker.currentStep != -1) &&
			(earliestWorker == -1 || worker.timeComplete < workers[earliestWorker].timeComplete) {
			earliestWorker = i
		}
	}
	if earliestWorker == -1 {
		return nil
	} else {
		return &workers[earliestWorker]
	}
}

// find the next step to execute and return that, or (-1, false) if no step is available
func findNextStep(input map[int][]int, numSteps int, startedSteps map[int]bool, completedSteps map[int]bool) (int, bool) {
	for i := 0; i < numSteps; i += 1 {
		if !startedSteps[i] {
			canDo := true
			for _, d := range input[i] {
				if !completedSteps[d] {
					canDo = false
				}
			}
			if canDo {
				return i, true
			}
		}
	}
	return -1, false
}

func stepWeight(step int) int {
	return 60 + 1 + step
}

func computeSteps(numWorkers int) ([]int, int) {
	input, numSteps := getInput()

	var stepOrder []int
	startedSteps := make(map[int]bool)
	completedSteps := make(map[int]bool)

	var workers []worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, worker { currentStep: -1, timeComplete: 0 })
	}

	timeNow := 0

	for {
		// find the next worker that will become available
		nextWorker := findNextWorker(workers, false)

		// fast forward time and mark the step the worker was executing as complete
		if nextWorker.timeComplete > timeNow {
			timeNow = nextWorker.timeComplete
			if nextWorker.currentStep != -1 {
				completedSteps[nextWorker.currentStep] = true
				nextWorker.currentStep = -1
			}
		}

		if len(completedSteps) == numSteps {
			return stepOrder, timeNow
		}

		// find next available step
		nextStep, nextStepAvailable := findNextStep(input, numSteps, startedSteps, completedSteps)

		if nextStepAvailable {
			// set the worker executing the next step
			stepOrder = append(stepOrder, nextStep)
			startedSteps[nextStep] = true
			nextWorker.currentStep = nextStep
			nextWorker.timeComplete = timeNow + stepWeight(nextStep)

		} else {
			// fast forward until the next worker finishes executing
			nextWorker := findNextWorker(workers, true)
			if nextWorker == nil {
				log.Fatal("No steps to do and no workers executing")
			}
			timeNow = nextWorker.timeComplete
			if nextWorker.currentStep != -1 {
				completedSteps[nextWorker.currentStep] = true
				nextWorker.currentStep = -1
			}
		}
	}
}

func part1() {
	steps, _ := computeSteps(1)
	var str strings.Builder
	for _, i := range steps {
		str.WriteString(string('A' + i))
	}
	fmt.Printf("part 1 result = %s\n", str.String())
}

func part2() {
	_, totalTime := computeSteps(5)
	fmt.Printf("part 2 result = %d\n", totalTime)
}

func Run() {
	part1()
	part2()
}
