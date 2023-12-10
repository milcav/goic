package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func swapMinGame(gameSlice []string, redval int, greenval int, blueval int) (int, int, int) {
	for _, value := range gameSlice {
		value = strings.Trim(value, " ")
		color := strings.Split(value, " ")[1]
		numb, err := strconv.Atoi(strings.Split(value, " ")[0])
		check(err)
		switch color {
		case "red":
			if numb > redval {
				redval = numb
			}
		case "green":
			if numb > greenval {
				greenval = numb
			}
		case "blue":
			if numb > blueval {
				blueval = numb
			}
		default:
			continue
		}
	}
	return redval, greenval, blueval
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	gameMap := make(map[string][]string)

	for fileScanner.Scan() {
		gameName := strings.Split(fileScanner.Text(), ":")[0]
		games := strings.Split(fileScanner.Text(), ":")[1]
		// fmt.Println("1. " + fileScanner.Text())
		// fmt.Println("2. " + gameName)
		// fmt.Println("3. " + games)
		games = strings.Trim(games, " ")
		gameList := strings.Split(games, ";")
		gameMap[gameName] = gameList
	}

	sumOfPowers := 0

	for key, value := range gameMap {
		redval := 0
		greenval := 0
		blueval := 0
		for _, gameString := range value {
			redval, greenval, blueval = swapMinGame(strings.Split(gameString, ", "), redval, greenval, blueval)
		}
		fmt.Println(value)
		fmt.Println(key, redval, greenval, blueval)
		sumOfPowers += redval * greenval * blueval
	}

	fmt.Println(sumOfPowers)
}
