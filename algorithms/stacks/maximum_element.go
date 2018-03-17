package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

//https://www.hackerrank.com/challenges/maximum-element/problem
type Instruction struct {
	instructionType  int64
	instructionValue int64
}

func main() {
	var stack []int64
	scanner := getScanner()
	for scanner.Scan() {
		instruction := getNextInstruction(strings.Split(scanner.Text(), " "))
		stack = parseNextInstruction(instruction, stack)
	}
}

func parseNextInstruction(instruction Instruction, stack []int64) []int64 {
	switch instruction.instructionType {
	default:
	case 0:
		break
	case 1:
		stack = append(stack, instruction.instructionValue)
	case 2:
		stack = stack[0:len(stack)-1]
	case 3:
		fmt.Println(MaxValue(stack))
	}

	return stack
}

func getScanner() *bufio.Scanner {
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanLines)
	return reader
}

func getNextInstruction(input []string) Instruction {
	list := MapStringToInt(input)
	return Instruction{
		readInstruction(list),
		readValue(list),
	}
}

func readInstruction(input []int64) int64 {
	if input[0] > 3 {
		return 0
	}

	return input[0]
}

func readValue(input []int64) int64 {
	if len(input) == 1 {
		return 0
	}

	return input[1]
}

func MapStringToInt(input []string) []int64 {
	var output = make([]int64, len(input))
	for index, value := range input {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			output[index] = parsed
		}
	}

	return output
}

func MaxValue(input []int64) int64 {
	var max int64 = 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}
