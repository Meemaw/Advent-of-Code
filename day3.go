package advent_of_code

import (
	"regexp"
	"strconv"
)

type Triangle struct {
	a int
	b int
	c int
}

func(tri *Triangle) isPossible() bool {
	return tri.a + tri.b > tri.c && tri.a + tri.c > tri.b && tri.b + tri.c > tri.a
}

func NumberOfPossibleTriangles(triangles []*Triangle) int {
	count := 0
	for _, triangle := range triangles {
		if(triangle.isPossible()) {
			count++
		}
	}
	return count
}

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

func parseInput(inputText string) [][]int {
	lines := RegSplit(inputText, "\n")
	numLines := len(lines) - 1
	lineList := make([][]int, numLines, numLines)

	for i := 0; i < numLines; i++ {
		split := RegSplit(lines[i], " +")
        lineA, _ := strconv.Atoi(split[1])
        lineB, _ := strconv.Atoi(split[2])
        lineC, _ := strconv.Atoi(split[3])
        lineList[i] = []int{lineA, lineB, lineC}
	}
	return lineList
}

func ExtractTriangles2(inputText string) []*Triangle {
	lineList := parseInput(inputText)
	numLines := len(lineList)
	triangles := make([]*Triangle, numLines, numLines)

	ix := 0
	for j := 0; j < 3; j++ {
		for i := 0; i < numLines; i += 3 {
			triangles[ix] = &Triangle{a : lineList[i][j], b : lineList[i+1][j], c : lineList[i+2][j]}
			ix++
		}
	}
	return triangles
}


func ExtractTriangles1(inputText string) []*Triangle {
	lineList := parseInput(inputText)
	numLines := len(lineList)
	triangles := make([]*Triangle, numLines, numLines)

	for i, line := range lineList {
		triangles[i] = &Triangle{a : line[0], b : line[1], c : line[2]}
	}
	return triangles
}
