package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Quiz struct {
	Ques string
	Ans  string
}

func main() {

	lines, err := ReadCsv("questions.csv")
	if err != nil {
		panic(err)
	}
	count := 0

	// Loop through lines & turn into object
	for i, line := range lines {
		data := Quiz{
			Ques: line[0],
			Ans:  line[1],
		}
		fmt.Println("Ques", i+1, ":", data.Ques)
		var a string
		fmt.Scanln(&a)
		if a == data.Ans {
			count++
		}
	}
	fmt.Println("Score is ", count, "/12")
	fmt.Println(Pass(count))
}

func Pass(n int) (string, float64) {
	p := (float64(n) / 12) * 100
	if p > 80 {
		return "Pass. Percentage :", p
	} else {
		return "Fail. Percentage :", p
	}
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
