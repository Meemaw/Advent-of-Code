package advent_of_code

import (
    "testing"
    "io/ioutil"
)

func TestDay3(t *testing.T) {

    bytes, _ := ioutil.ReadFile("day3_input.txt")
    inputText := string(bytes)

    // part 1
    triangles1 := ExtractTriangles1(inputText)
    count1 := NumberOfPossibleTriangles(triangles1)

    if count1 != 983 {
        t.Errorf("Got %v expected %v", count1, 983)
    }

    // part 2
    triangles2 := ExtractTriangles2(inputText)
    count2 := NumberOfPossibleTriangles(triangles2)
    if count2 != 1836 {
        t.Errorf("Got %v expected %v", count2, 1836)
    }
}
