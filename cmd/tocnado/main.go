package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Headline struct {
	text        string
	level       int
	chainedText string
}

func newHeadline(text string, level int) Headline {
	headline := Headline{
		text:        text,
		level:       level,
		chainedText: strings.ToLower(strings.Replace(text, " ", "-", -1)),
	}
	return headline
}

type TableOfContent struct {
	headlines []Headline
}

func main() {
	file, err := os.Open("test.md")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	headlinePattern := `(#+)\s(.*)`
	headlineRegex := regexp.MustCompile(headlinePattern)

	scanner := bufio.NewScanner(file)
	var headlines []Headline
	for scanner.Scan() {
		line := scanner.Text()
		if headlineRegex.MatchString(line) {
			match := headlineRegex.FindStringSubmatch(line)
			headline := newHeadline(match[2], len(match[1]))
			headlines = append(headlines, headline)
		}
	}

	fmt.Println(headlines)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
