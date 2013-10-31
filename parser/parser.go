package parser

import (
	"regexp"
	"strconv"
	"strings"
)

func MemorableCharacters(body []byte, memorableCharacters string) [3]string {
	memSplit := strings.Split(memorableCharacters, "")

	bodyStr := string(body)

	characters := []string{
		regexp.MustCompile(`memInfo1\">Character (\d+)`).FindStringSubmatch(bodyStr)[1],
		regexp.MustCompile(`memInfo2\">Character (\d+)`).FindStringSubmatch(bodyStr)[1],
		regexp.MustCompile(`memInfo3\">Character (\d+)`).FindStringSubmatch(bodyStr)[1],
	}

	var characterIndexes [3]int

	for i, char := range characters {
		val, _ := strconv.Atoi(char)
		characterIndexes[i] = val - 1
	}

	return [3]string{
		memSplit[characterIndexes[0]],
		memSplit[characterIndexes[1]],
		memSplit[characterIndexes[2]],
	}
}

func SubmitToken(body []byte) string {
	submitRegex := regexp.MustCompile(`name="submitToken" value="(\d+)`)
	return submitRegex.FindStringSubmatch(string(body))[1]
}
