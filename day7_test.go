package advent_of_code

import (
    "testing"
    "io/ioutil"
)

func TestDay7(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day7_input.txt")
    inputText := string(bytes)

	if code := NumSupportingIps(inputText, true, false); code != 118 {
		t.Errorf("Got %v expected %v", code, 118)
	}

	if code := NumSupportingIps(inputText, false, true); code != 260 {
		t.Errorf("Got %v expected %v", code, 260)
	}
}
