package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	var onlyNumberLines []string
	re := regexp.MustCompile(`\D`)
	sum := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		res := re.ReplaceAllString(str, "")
		onlyNumberLines = append(onlyNumberLines, res)
	}

	for _, number := range onlyNumberLines {

		real_number, err := strconv.Atoi((string(number[0]) + string(number[len(number)-1:])))
		check(err)
		fmt.Printf("This is the string number: %s\n", number)
		fmt.Printf("This is the number: %d\n", real_number)
		sum += real_number
	}
	dat.Close()
	fmt.Println(sum)
}
