package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	var matrix [][]string
	// numbersReg := regexp.MustCompile(`[0-9]+`)
	// signsReg := regexp.MustCompile(`[^\w|\d|\.]`)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		// res := signsReg.FindAllStringIndex(str, -1)
		// fmt.Println(res)
		// fmt.Println("\n")
		sli := strings.Split(str, "")
		matrix = append(matrix, sli)
	}

	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			fmt.Printf("%s ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
