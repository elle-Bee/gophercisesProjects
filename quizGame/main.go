package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// CSV file parsing
	csvFilename := flag.String("csv", "problems.csv", "CSV file in format of 'question,answer' to be importeed")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErrNil(err)

	r := csv.NewReader(file)
	listProblems, err := r.ReadAll()
	checkErrNil(err)

	defer file.Close()

	//Questionaire
	var correct int
	for i, p := range parseProblems(listProblems) {
		var ans string
		fmt.Printf("Q%d: %s = ", i+1, p.q)
		fmt.Scanf("%s", &ans)
		if ans == p.a {
			fmt.Printf("Correct! \n")
			correct++
		} else {
			fmt.Printf("Oops! That's incorrect. \n")
		}
	}
	fmt.Printf("Congratulations! You've scored %d/%d", correct, len(listProblems))

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}

type problem struct {
	q string
	a string
}

func parseProblems(listProblems [][]string) []problem {
	returnedProblems := make([]problem, len(listProblems))
	for i, line := range listProblems {
		returnedProblems[i] = problem{
			// Validates the CSV file so that no spaces are counted
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return returnedProblems
}
