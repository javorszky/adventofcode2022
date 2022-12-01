package day1

import (
	"fmt"
	"log"

	"github.com/javorszky/adventofcode2022/inputs"
)

func Task1() {
	gog, err := inputs.GroupByBlankLines("day1/input1_example.txt")
	if err != nil {
		log.Fatalf("day 1 task 1 could not read info: %s", err.Error())
	}
	fmt.Printf("%#v", gog)
}
