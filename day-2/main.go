package main

import (
	"aoc/day2/submarine"
	"fmt"
	"log"
)

func main() {
	s := submarine.Submarine{HorizontalPosition: 0, Depth: 0}
	instructionList, err := submarine.ReadInstructionFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range instructionList {
		switch instruction.Command {
		case "forward":
			s.MoveForward(instruction.GridCount)
		case "down":
			s.Dive(instruction.GridCount)
		case "up":
			s.Rise(instruction.GridCount)
		}
	}
	position := s.Depth * s.HorizontalPosition
	fmt.Printf("submarine Data: %v\n", s)
	fmt.Printf("Position: %d", position)
}
