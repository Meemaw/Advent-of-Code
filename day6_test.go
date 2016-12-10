package advent_of_code

import (
    "testing"
    "io/ioutil"
)

func TestDay6(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day6_input.txt")
    inputText := string(bytes)

	if code := GetErrorCorrectedMessage(inputText, true); code != "ygjzvzib" {
		t.Errorf("Got %v expected %v", code, "ygjzvzib")
	}

	if code := GetErrorCorrectedMessage(inputText, false); code != "pdesmnoz" {
		t.Errorf("Got %v expected %v", code, "pdesmnoz")
	}
}
