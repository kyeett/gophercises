package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\"")
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

	scanner := bufio.NewScanner(os.Stdin)

	score := 0
	for i, record := range records {
		var answer string

		question := record[0]
		correctAnswer := record[1]

		fmt.Printf("Problem #%d: %s = ", i+1, question)

		//Get answer
		scanner.Scan()
		answer = scanner.Text()
		if answer == "q" {
			break
		}
		//Check if answer is correct
		if strings.Compare(answer, correctAnswer) == 0 {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, len(records))

	/*	"-csv string"
		"-limit int"


	*/

}
