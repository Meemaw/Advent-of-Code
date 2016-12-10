package advent_of_code

import(
	"sort"
	"strings"
	"strconv"
)

const (
	NumLetters = 26
	LowestAscii = 97
	HighestAscii = 122
)

func IsRealRoom(encrypted string) int {
	if(encrypted == "") {
		return -1
	}
	occurancesList := make([]CharOccurances, 256, 256)

	checksumSplit := RegSplit(encrypted, "-[0-9]+")
	s1 := strings.Split(encrypted, "[")
	s2 := strings.Split(s1[0], "-")
	sectorId, _ := strconv.Atoi(s2[len(s2)-1])
	message := checksumSplit[0]
	checksum := checksumSplit[1][1:len(checksumSplit[1])-1]

	messageSplit := RegSplit(message, "-")

	for _, sector := range messageSplit {
		for _, char := range sector {
			occurancesList[char].value = char
			occurancesList[char].numOccurances++
		}
	}

	// sort the list by Occurances and lexicographical
	sort.Sort(ByOccurances(occurancesList))

	for i, char := range checksum {
		if(occurancesList[i].value != char) {
			return -1
		}
	}
	return sectorId
}

func NumberOfRealRooms(inputString string) int {
	count := 0
	rooms := RegSplit(inputString, "\n")
	for _, room := range rooms {
		sectorId := IsRealRoom(room)
		if(sectorId >= 0) {
			count += sectorId
		}
	}
	return count
}

func rotateChar(char rune, numTimes int) byte {
	actualNumTimes := numTimes % NumLetters
	newChar := int(char) + actualNumTimes
	if(newChar > HighestAscii) {
		diff := newChar - HighestAscii
		newChar = LowestAscii + diff - 1
	}

	return byte(newChar)
}

func rotateWord(word string, sectorId int) string {
	s := ""
	for _, char := range word {
		newChar := rotateChar(char, sectorId)
		s += string(newChar)
	}
	return s
}


func DecodeRealRooms(inputString string) *List {
	decodedRealRooms := NewList()

	rooms := RegSplit(inputString, "\n")

	for _, room := range rooms {
		sectorId := IsRealRoom(room)

		// if is real room
		if(sectorId > 0) {
			code := RegSplit(room, "-[0-9]+")
    		split := strings.Split(code[0], "-")
    		for i, word := range split {
    			split[i] = rotateWord(word, sectorId)

    		}
    		decodedRealRooms.Add(strings.Join(split, " ") + " " + strconv.Itoa(sectorId))
		}
	}
	return decodedRealRooms
}
