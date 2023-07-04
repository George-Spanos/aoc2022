package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	calories, index int
}

func main() {
	topElves := make([]Elf, 3)
	var idx int
	var lastElf Elf
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			lastElf.index = idx
			belongToTopElves(lastElf, topElves)
			idx += 1
			lastElf.calories = 0
		} else {
			cals, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			lastElf.calories += cals
		}
	}
	var totalCalories int
	for _, e := range topElves {
		totalCalories += e.calories
	}
	fmt.Printf("Top 3 Elves with the most calories have a total %v\n", totalCalories)
}
func belongToTopElves(elf Elf, topElves []Elf) {
	smallerTopElf := &topElves[0]
	var idx int
	for i := range topElves {
		if smallerTopElf.calories > topElves[i].calories {
			smallerTopElf = &topElves[i]
			idx = i
		}
	}
	if elf.calories > smallerTopElf.calories {
		topElves[idx] = elf
	}
}
