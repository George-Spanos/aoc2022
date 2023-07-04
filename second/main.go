package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Round struct {
	Opponent string
	Player   string
}

var score int

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		selection := pickSelection(args[0], args[1])
		fmt.Println(args, "selection", selection, "victoryScore", victoryScore(args[0], selection), "selection score", extractSectionScore(selection))
		score += victoryScore(args[0], selection) + extractSectionScore(selection)
	}
	log.Println(score)
}
func pickSelection(opponent, result string) string {
	switch result {
	case "X":
		switch opponent {
		case "A":
			return "C"
		case "B":
			return "A"
		case "C":
			return "B"
		}
	case "Y":
		return opponent
	case "Z":
		switch opponent {
		case "A":
			return "B"
		case "B":
			return "C"
		case "C":
			return "A"
		}
	}
	return ""
}
func extractSectionScore(selection string) int {
	switch selection {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	default:
		return 0
	}
}
func victoryScore(opponent, player string) int {
	if opponent == player {
		return 3
	}
	if player == "A" && opponent == "C" || player == "B" && opponent == "A" || player == "C" && opponent == "B" {
		return 6
	}
	return 0
}
