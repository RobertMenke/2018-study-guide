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
	instructionType  int32
	instructionValue int32
}

func main() {
	var stack []int32
	scanner := getScanner()
	for scanner.Scan() {
		instruction := getNextInstruction(strings.Split(scanner.Text(), " "))
		stack = parseNextInstruction(instruction, stack)
	}
}

func parseNextInstruction(instruction Instruction, stack []int32) []int32 {
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

func readInstruction(input []int32) int32 {
	if input[0] > 3 {
		return 0
	}

	return input[0]
}

func readValue(input []int32) int32 {
	if len(input) == 1 {
		return 0
	}

	return input[1]
}

func MapStringToInt(input []string) []int32 {
	var output = make([]int32, len(input))
	for index, value := range input {
		parsed, err := strconv.ParseInt(value, 10, 32)
		if err == nil {
			output[index] = int32(parsed)
		}
	}

	return output
}

func MaxValue(input []int32) int32 {
	var max int32 = 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}
