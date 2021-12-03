package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	instruction string
	value       int
}

func main() {

	var instructionList []Instruction
	var horizontalPosition int = 0
	var depth int = 0
	var aim int = 0

	fmt.Println("Hello!")

	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instructionList = append(instructionList, NewInstruction(scanner.Text()))
	}

	for _, element := range instructionList {
		fmt.Println(element)
		switch element.instruction {
		case "forward":
			horizontalPosition += element.value
			depth += aim * element.value
		case "up":
			aim -= element.value
		case "down":
			aim += element.value
		default:

		}
	}

	fmt.Printf("Depth: %d Horizontal Pos: %d, Aim: %d\n", depth, horizontalPosition, aim)
	fmt.Printf("%d", depth*horizontalPosition)
}

func NewInstruction(instruction string) Instruction {
	var splitInstruction = strings.Split(instruction, " ")[0]

	var value, _ = strconv.Atoi(strings.Split(instruction, " ")[1])

	return Instruction{splitInstruction, value}
}
