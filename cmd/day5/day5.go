package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"strconv"
	"strings"
)

type Ingredient int

type Range struct {
	first Ingredient
	last  Ingredient
}

func stringToIngredient(str string) Ingredient {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return Ingredient(i)
}

func parseInput(input string) (ranges []Range, ingredients []Ingredient) {

	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		panic(input)
	}

	line_iter := strings.SplitSeq(parts[0], "\n")
	for line := range line_iter {
		FirstLast := strings.Split(line, "-")
		if len(FirstLast) != 2 {
			panic(line)
		}
		first := stringToIngredient(FirstLast[0])
		last := stringToIngredient(FirstLast[1])
		ranges = append(ranges, Range{first, last})
	}

	line_iter = strings.SplitSeq(parts[1], "\n")
	for line := range line_iter {
		ingredient := stringToIngredient(line)
		ingredients = append(ingredients, ingredient)
	}
	return
}

func isFresh(ingredient Ingredient, ranges []Range) bool {
	for _, r := range ranges {
		if ingredient >= r.first && ingredient <= r.last {
			return true
		}
	}
	return false
}

func countFresh(ingredients []Ingredient, ranges []Range) (counter int) {
	for _, ingredient := range ingredients {
		if isFresh(ingredient, ranges) {
			counter++
		}
	}
	return
}

func main() {
	ranges, ingredients := parseInput(util.LoadInput(5))
	counter := countFresh(ingredients, ranges)
	fmt.Println(counter)
}
