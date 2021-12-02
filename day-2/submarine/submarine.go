package submarine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	HorizontalPosition int
	Depth              int
	Aim                int
}

type Instruction struct {
	Command   string
	GridCount int
}

func ReadInstructionFile(filePath string) ([]Instruction, error) {
	instructionList := []Instruction{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read instruction file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		instructionSet := strings.Split(fileScanner.Text(), " ")
		command := instructionSet[0]
		gridCount, err := strconv.Atoi(instructionSet[1])
		if err != nil {
			return nil, fmt.Errorf("cannot parse gridCount in instruction file: %s", err)
		}
		instructionList = append(instructionList, Instruction{Command: command, GridCount: gridCount})
	}

	if err := fileScanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading instruction file: %s", err)
	}

	return instructionList, nil

}

func (s *Submarine) MoveForward(gridCount int) {
	s.HorizontalPosition += gridCount
	s.Depth += s.Aim * gridCount
}

func (s *Submarine) Dive(gridCount int) {
	s.Aim += gridCount
}

func (s *Submarine) Rise(gridCount int) {
	s.Aim -= gridCount
}
