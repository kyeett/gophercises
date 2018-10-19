package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\"")
	limitPtr := flag.Int("limit", 10, "the time limit for the quiz in seconds (default 10)")
	flag.Parse()

	f, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	fr := csv.NewReader(f)

	records, err := fr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	questions := parseRecords(records)

	score := runQuiz(questions, *limitPtr)
	fmt.Printf("You scored %d out of %d\n", score, len(questions))
}

func runQuiz(questions []question, limit int) (score int) {
	answerCh := make(chan string)
	// Goroutine for user input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			//Get answer
			scanner.Scan()
			answerCh <- scanner.Text()
		}
	}()

	// Accept answers until timeout
	timeout := time.NewTimer(time.Duration(limit) * time.Second)
	for i, question := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, question.q)
		select {
		case <-timeout.C:
			fmt.Printf("\nTimeout after %d second(s)\n", limit)
			return
		case answer := <-answerCh:
			if strings.Compare(answer, question.a) == 0 {
				score++
			}
		}
	}
	return
}

func parseRecords(records [][]string) []question {
	var ret []question
	for _, record := range records {
		ret = append(ret, question{
			q: record[0],
			a: record[1],
		})
	}
	return ret
}

type question struct {
	q string
	a string
}
