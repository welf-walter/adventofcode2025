package main

import (
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
