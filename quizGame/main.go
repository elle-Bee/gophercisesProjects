package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "CSV file in format of 'question,answer' to be importeed")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErrNil(err)

	r := csv.NewReader(file)
	listProblems, err := r.ReadAll()
	checkErrNil(err)
	fmt.Println(listProblems)

	defer file.Close()

	parseProblems(listProblems)

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
			q: line[0],
			a: line[1],
		}
	}
	return returnedProblems
}
