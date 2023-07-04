package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func b() {
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
	i := 0
	elves := make([]string, 6)
	for scanner.Scan() {
		if i <= 5 {
			elves[i] = scanner.Text()
			if i != 5 {
				i++
				continue
			} else {
				i = 0
			}
		}
		duplicate, found := findDuplicate(elves[0:3]...)
		if !found {
			log.Fatal("no duplicate found")
		}
		sum += priorities[duplicate]

		duplicate, found = findDuplicate(elves[3:6]...)
		if !found {
			log.Fatal("no duplicate found")
		}
		sum += priorities[duplicate]
	}
	fmt.Println("Priority Sum of badges", sum)
}
