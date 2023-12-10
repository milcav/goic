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

func checkGame(gameSlice []string) bool {
	var ret bool = true
	for _, value := range gameSlice {
		value = strings.Trim(value, " ")
		color := strings.Split(value, " ")[1]
		numb, err := strconv.Atoi(strings.Split(value, " ")[0])
		check(err)
		switch color {
		case "red":
			if numb > 12 {
				ret = false
			}
		case "green":
			if numb > 13 {
				ret = false
			}
		case "blue":
			if numb > 14 {
				ret = false
			}
		default:
			ret = true
		}
	}
	return ret
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

	sumOfGames := 0

	for key, value := range gameMap {
		ret := true
		for _, gameString := range value {
			if !checkGame(strings.Split(gameString, ", ")) {
				ret = false
			}
		}
		if ret {
			key = strings.Trim(key, " ")
			numb, err := strconv.Atoi(strings.Split(key, " ")[1])
			check(err)
			sumOfGames += numb
		}
	}

	fmt.Println(sumOfGames)
}
