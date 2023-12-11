package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// var onlyNumberLines []string
	// numbersReg := regexp.MustCompile(`[0-9]+`)
	signsReg := regexp.MustCompile(`[^\w|\d|\.]`)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		res := signsReg.FindAllStringIndex(str, -1)
		fmt.Println(res)
		fmt.Println("\n")
		// onlyNumberLines = append(onlyNumberLines, res)
	}

	// for _, number := range onlyNumberLines {
	// 	fmt.Println(number)
	// }
	dat.Close()
}
