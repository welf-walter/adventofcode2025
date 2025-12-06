package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"log"
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

func parseInput2(input string) (problems []problem) {
	lines := strings.Split(input, "\n")
	h := len(lines) - 1
	w := len(lines[0])
	for _, line := range lines {
		w = max(w, len(line))
	}

	operatorLine := lines[h]
	tokenIter := strings.SplitSeq(operatorLine, " ")
	for token := range tokenIter {
		switch token {
		case "":
		case string(ADD):
			problems = append(problems, problem{operator: ADD})
		case string(MULTIPLY):
			problems = append(problems, problem{operator: MULTIPLY})
		default:
			panic(token)
		}
	}
	log.Println(problems)

	p := len(problems) - 1
	for x := w - 1; x >= 0; x-- {
		number := 0
		for y := 0; y < h; y++ {
			if x < len(lines[y]) {
				c := lines[y][x]
				digit, err := strconv.Atoi(string(c))
				if err == nil {
					number = number*10 + digit
				}
			}
		}
		if number > 0 {
			log.Println(number)
			problems[p].operands = append(problems[p].operands, number)
		} else {
			p--
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

func sumSolvedProblems(problems []problem) (sum int) {
	for _, p := range problems {
		sum += solveProblem(p)
	}
	return
}

func main() {
	problems := parseInput(util.LoadInput(6))
	fmt.Println(sumSolvedProblems(problems))
}
