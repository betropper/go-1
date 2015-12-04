package main

import (
	"fmt"
	"github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	"github.com/skilstak/go/input"
	"os"
	"strings"
)

var answers = []string{"YEEEEEEEEESSSSS.", "NAHHHHHHHHH.", "UH. I MEAN. MAYBE?", "GEEZ, I DUNNO, STOP ASKING ME SO MANY QUESTIONS."}

func bye() {
	fmt.Println("")
	fmt.Println("WELL OKAY BYE I GUESS")
	os.Exit(-1)
}

func main() {
	fmt.Println(c.CL + c.R + "MAGIC EIGHT BALLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.")
	fmt.Println(c.X + c.R + "GIMME A QUESTION OR SOMETHING.")
	for true {
		question := input.Ask(c.X + c.R + "> " + c.B3)
		if strings.Contains(question, "love") {
			fmt.Println(c.R + "u w0t m8. Srsly? Love? Stahp.")
		} else if strings.Contains(question, "bye") {
			bye()
		} else {
			answer := choice.Strings(answers)
			fmt.Println(c.Random() + answer + c.X)
		}
	}
}
