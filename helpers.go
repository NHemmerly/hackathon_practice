package main

import (
	"fmt"
	"strings"
)

type StatTracker struct {
	WordList       map[string]int
	Unique         int
	TotalWords     int
	N5Counts       int
	N4Counts       int
	N3Counts       int
	N2Counts       int
	N1Counts       int
	UngradedCounts int
}

func (s StatTracker) buildHtmxOut(tokens *tokens) (string, error) {
	var sBuilder strings.Builder
	sBuilder.WriteString("<p>")
	for _, token := range *tokens {
		s.countUniqueWords(token.Lem, token.Pos)
		var bgColor string
		switch token.Difficulty {
		case "1":
			bgColor = "bg-red-400"
		case "2":
			bgColor = "bg-orange-300"
		case "3":
			bgColor = "bg-amber-200"
		case "4":
			bgColor = "bg-green-300"
		case "5":
			bgColor = "bg-sky-300"
		}
		htmx_out := fmt.Sprintf("<span class=\"max-w-xl %s\">%s</span>", bgColor, token.Surface)
		sBuilder.WriteString(htmx_out)
	}
	fmt.Println(s.WordList)
	sBuilder.WriteString("</p>")
	return sBuilder.String(), nil
}

func (s StatTracker) countUniqueWords(word string, pos string) {
	if pos == "補助記号-句点" || pos == "補助記号-読点" || pos == "補助記号-括弧開" {
		return
	}
	_, ok := s.WordList[word]
	if !ok {
		s.WordList[word] = 1
	} else {
		s.WordList[word] += 1
	}
}
