package solution

import (
	"fmt"
	"sort"
)

func CountAndSortAppearance(elems []int) []PairContainer {
	pairs := countAppearance(elems)
	sortAppearance(pairs)
	return pairs
}

func countAppearance(elems []int) []PairContainer {
	counter := make(map[int]int)
	for _, elem := range elems {
		counter[elem] += 1
	}

	appearance := make([]PairContainer, 0)
	for elem, count := range counter {
		appearance = append(appearance, PairContainer{
			Elem:  elem,
			Count: count,
		})
	}

	return appearance
}

type PairContainer struct {
	Elem  int
	Count int
}

func sortAppearance(pairs []PairContainer) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})
}

func PrintAppearance(pairs []PairContainer) {
	fmt.Println("number --> total")
	for _, pair := range pairs {
		fmt.Printf("%d --> %d\n", pair.Elem, pair.Count)
	}
}
