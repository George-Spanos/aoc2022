package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func a() {
	priorities := make(map[rune]int)
	for i, c := 1, 'a'; c <= 'z'; {
		priorities[c] = i
		i++
		c++
	}
	for i, c := 27, 'A'; c <= 'Z'; {
		priorities[c] = i
		i++
		c++
	}
	file, err := os.Open("./input")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		rubersack := scanner.Text()
		comp1, comp2 := readCompartments(rubersack)
		duplicate, found := findDuplicate(comp1, comp2)
		if found {
			sum += priorities[duplicate]
		} else {
			log.Fatal("did not have a duplicate")
		}
	}
	fmt.Println("Priority Sum of duplicates", sum)
}
