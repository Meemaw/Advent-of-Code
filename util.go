package advent_of_code

import(
	"strings"
	"regexp"
)

// interface for sorting chars by occurances
type ByOccurances []CharOccurances

func (slice ByOccurances) Len() int {
	return len(slice)
}

func (slice ByOccurances) Swap(i,j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice ByOccurances) Less(i,j int) bool {
	if(slice[i].numOccurances == slice[j].numOccurances) {
		return slice[i].value < slice[j].value
	} else {
		return slice[i].numOccurances > slice[j].numOccurances
	}
}

func GetLines(inputString string) []string {
	split := strings.Split(inputString, "\n")
	return split[0:len(split)-1]
}
// interface for sorting chars by occurances



// splitting on regex
func RegSplit(text string, delimeter string) []string {
    reg := regexp.MustCompile(delimeter)
    indexes := reg.FindAllStringIndex(text, -1)
    laststart := 0
    result := make([]string, len(indexes) + 1)
    for i, element := range indexes {
            result[i] = text[laststart:element[0]]
            laststart = element[1]
    }
    result[len(indexes)] = text[laststart:len(text)]
    return result
}
