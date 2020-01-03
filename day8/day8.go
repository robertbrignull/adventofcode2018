package day8

import (
	"fmt"
	"github.com/robertbrignull/adventofcode2018/util"
	"strconv"
	"strings"
)

type licenseNode struct {
	children []*licenseNode
	metadata []int
}

func parseLicenseNode(input []string) (*licenseNode, []string) {
	numChildren, _ := strconv.Atoi(input[0])
	numMetadata, _ := strconv.Atoi(input[1])
	input = input[2:]

	var children []*licenseNode
	for i := 0; i < numChildren; i++ {
		var child *licenseNode
		child, input = parseLicenseNode(input)
		children = append(children, child)
	}

	var metadata []int
	for i := 0; i < numMetadata; i++ {
		x, _ := strconv.Atoi(input[i])
		metadata = append(metadata, x)
	}
	input = input[numMetadata:]

	node := licenseNode { children, metadata }
	return &node, input
}

func getInput() *licenseNode {
	input := strings.Split(util.ReadFile("./day8/day8_input"), " ")
	node, _ := parseLicenseNode(input)
	return node
}

func totalMetadata(node *licenseNode) int {
	total := 0
	for i := 0; i < len(node.children); i++ {
		total += totalMetadata(node.children[i])
	}
	for i := 0; i < len(node.metadata); i++ {
		total += node.metadata[i]
	}
	return total
}

func nodeValue(node *licenseNode) int {
	value := 0
	if len(node.children) == 0 {
		for i := 0; i < len(node.metadata); i++ {
			value += node.metadata[i]
		}
	} else {
		for i := 0; i < len(node.metadata); i++ {
			child := node.metadata[i] - 1
			if child >= 0 && child < len(node.children) {
				value += nodeValue(node.children[child])
			}
		}
	}
	return value
}

func Run() {
	rootNode := getInput()
	fmt.Printf("part 1 result = %d\n", totalMetadata(rootNode))
	fmt.Printf("part 2 result = %d\n", nodeValue(rootNode))
}
