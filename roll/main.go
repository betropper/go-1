package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"github.com/skilstak/go/input"
	"math/rand"
	"os"
	"strconv"
)

func rolldice(number string) int {
	n, _ := strconv.Atoi(number)
	total := 0
	for i := 0; i <= n; i++ {
		oneroll := rand.Intn(6)
		printedroll := strconv.Itoa(oneroll)
		print(c.X + c.Random() + printedroll + " ")
		total += oneroll
	}
	fmt.Println("")
	return total
}

func main() {
	var dicecount string
	if len(os.Args) == 2 {
		dicecount = os.Args[1]
	} else {
		fmt.Println(c.Clear + c.Cyan + "How many dice would you like to roll?")
		dicecount = input.Ask(c.X + c.Magenta + "")
	}
	printedtotal := strconv.Itoa(rolldice(dicecount))
	fmt.Println(c.X + c.Cyan + "You rolled " + c.Magenta + dicecount + c.X + c.Cyan + " dice, for a total of " + c.X + c.Magenta + printedtotal + ".")
}
