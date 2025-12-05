package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"log"
	"slices"
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

// convert 3-7 and 6-9 to 3-9
// ok = false if not convertable
func unionRange(r1 Range, r2 Range) (r Range, ok bool) {
	// ***
	//      ***
	if r1.last+1 < r2.first {
		return Range{0, 0}, false
	}
	//      ***
	// ***
	if r1.first-1 > r2.last {
		return Range{0, 0}, false
	}
	// ********
	//    *******
	return Range{min(r1.first, r2.first), max(r1.last, r2.last)}, true
}

func condenseRanges(ranges []Range) (condensed []Range) {
	log.Printf("In: %v\n", ranges)
	cmp := func(a, b Range) int { return int(a.first) - int(b.first) }
	slices.SortFunc(ranges, cmp)

	log.Printf("Sorted: %v\n", ranges)
	for _, r := range ranges {
		if len(condensed) > 0 {
			last := condensed[len(condensed)-1]
			union, ok := unionRange(r, last)
			if ok {
				condensed[len(condensed)-1] = union
			} else {
				condensed = append(condensed, r)
			}

		} else {
			condensed = []Range{r}
		}
		log.Printf("   %v\n", condensed)
	}

	log.Printf("Out: %v\n", condensed)

	return
}

func sumRanges(ranges []Range) (sum int) {
	for _, r := range ranges {
		sum += int(r.last - r.first + 1)
	}
	return
}

func main() {
	ranges, ingredients := parseInput(util.LoadInput(5))
	counter := countFresh(ingredients, ranges)
	fmt.Println(counter)

	// part 2
	ranges = condenseRanges(ranges)
	fmt.Println(sumRanges(ranges))

}
