package main

import (
	"log"
)

func main() {
	a()
	b()
}
func readCompartments(rucksack string) (string, string) {
	if len(rucksack)%2 != 0 {
		log.Fatal("rucksack length is not even")
	}
	return rucksack[0 : len(rucksack)/2], rucksack[len(rucksack)/2:]
}
func findDuplicate(comps ...string) (rune, bool) {
	if len(comps) < 2 {
		return '0', false
	}
	// convert arrays to maps for easy look up
	maps := make([]map[rune]bool, len(comps))
	for i, s := range comps {
		m := make(map[rune]bool)
		for _, c := range s {
			m[c] = true
		}
		maps[i] = m
	}
	for _, c := range comps[0] {
		exists := true
		for i := range maps {
			if _, ok := maps[i][c]; !ok {
				exists = false
				break
			}
		}
		if exists {
			return c, exists
		}
	}
	return '0', false
}
