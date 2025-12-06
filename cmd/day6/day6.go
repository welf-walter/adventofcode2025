package main

import (
	"strconv"
	"strings"
)

const ADD = '+'
const MULTIPLY = '*'

type problem struct {
	operands []int
	operator byte
}

/*
const newProblem = problem{
	operands: nil,
	operator: '?',
}
*/

func parseInput(input string) (problems []problem) {
	linesIter := strings.SplitSeq(input, "\n")
	for line := range linesIter {
		p := 0
		tokenIter := strings.SplitSeq(line, " ")
		for token := range tokenIter {
			switch token {
			case "":
			case string(ADD):
				problems[p].operator = ADD
				p++
			case string(MULTIPLY):
				problems[p].operator = MULTIPLY
				p++
			default:
				{
					number, err := strconv.Atoi(token)
					if err != nil {
						panic(err)
					}
					if p >= len(problems) {
						problems = append(problems, problem{})
					}
					problems[p].operands = append(problems[p].operands, number)
				}
				p++
			}
		}
	}
	return
}

func solveProblem(problem problem) int {
	switch problem.operator {
	case ADD:
		{
			sum := 0
			for _, operand := range problem.operands {
				sum += operand
			}
			return sum
		}
	case MULTIPLY:
		{
			product := 1
			for _, operand := range problem.operands {
				product *= operand
			}
			return product
		}
	default:
		panic(problem.operator)
	}
}
