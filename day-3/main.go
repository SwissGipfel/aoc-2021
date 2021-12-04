package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	binSlice := getBinarySlice("./input.txt")
	counter := countPositionInSlice(binSlice)
	binRepresentation := []rune{}
	negatives := []rune{}
	for _, val := range counter {
		if val > 0 {
			binRepresentation = append(binRepresentation, '1')
			negatives = append(negatives, '0')
			continue
		}
		if val < 0 {
			binRepresentation = append(binRepresentation, '0')
			negatives = append(negatives, '1')
			continue
		} else {
			log.Fatal("whoops")
		}
	}
	gammaRateBinary := string(binRepresentation)
	epsilonRateBinary := string(negatives)

	gammaTransformed, _ := strconv.ParseInt(gammaRateBinary, 2, 64)
	epsilonTransformed, _ := strconv.ParseInt(epsilonRateBinary, 2, 64)

	powerconSumption := gammaTransformed * epsilonTransformed

	fmt.Println(powerconSumption)

}

func getBinarySlice(filepath string) []string {
	binarySlice := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	fscan := bufio.NewScanner(file)
	for fscan.Scan() {
		binarySlice = append(binarySlice, fscan.Text())
	}
	return binarySlice
}

func countPositionInSlice(s []string) []int {
	positionCounterLength := len(s[0])
	positionCounter := make([]int, positionCounterLength)
	for _, value := range s {
		for index, char := range value {
			if char == '1' {
				positionCounter[index]++
				continue
			}
			if char == '0' {
				positionCounter[index]--
			}
		}
	}
	return positionCounter
}
