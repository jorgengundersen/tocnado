package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Headline struct {
	rawText     string
	text        string
	level       int
	chainedText string
	headlineID  string
	anchorLink  string
}

func newHeadline(rawHeadline string) Headline {
	headlinePattern := `(#+)\s(.*)`
	headlineRegex := regexp.MustCompile(headlinePattern)

	headlineParts := headlineRegex.FindStringSubmatch(rawHeadline)

	headline := Headline{
		rawText: rawHeadline,
		text:    string(headlineParts[2]),
		level:   len(headlineParts[1]),
	}

	headline.chainedText = strings.ToLower(strings.Replace(headline.text, " ", "-", -1))
	headline.headlineID = "#" + headline.chainedText
	headline.anchorLink = fmt.Sprintf("[%s](%s)", headline.text, headline.headlineID)

	return headline
}

type TableOfContent struct {
	headlines []Headline
}

func getFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getHeadlines(lines []string) []Headline {
	headlinePattern := `(#+)\s(.*)`
	headlineRegex := regexp.MustCompile(headlinePattern)

	var headlines []Headline
	for _, line := range lines {
		if headlineRegex.MatchString(line) {
			headline := newHeadline(line)
			headlines = append(headlines, headline)
		}
	}
	return headlines
}

func main() {
	filePath := "test.md"
	lines, err := getFileLines(filePath)
	if err != nil {
		fmt.Println(err)
	}

	headlines := getHeadlines(lines)

	for _, headline := range headlines {
		fmt.Println(headline.anchorLink)
	}
}
