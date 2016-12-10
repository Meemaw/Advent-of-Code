package advent_of_code

import (
	"strings"
	"sort"
)

func getCorrectedCharacter(lines []string, column int, mostFrequent bool) rune {
	occurancesList := make([]CharOccurances, 256, 256)

	for _, word := range lines {
		occurancesList[word[column]].value = rune(word[column])
		occurancesList[word[column]].numOccurances++
	}

	sort.Sort(ByOccurances(occurancesList))
	if(mostFrequent) {
		return occurancesList[0].value
	} else {
		// find the first nonzero occurance character (least commong)
		for i:= 255; i >= 0; i-- {
			if(occurancesList[i].numOccurances > 0) {
				return occurancesList[i].value
			}
		}
	}
	return -1
}

func GetErrorCorrectedMessage(inputString string, mostFrequent bool) string {
	lines := strings.Split(inputString, "\n")
	lines = lines[0:len(lines)-1]
	numChars := len(lines[0])

	errorCorrectedMessage := ""
	for i := 0; i < numChars; i++ {
		errorCorrectedMessage += string(getCorrectedCharacter(lines, i, mostFrequent))
	}

	return errorCorrectedMessage
}
