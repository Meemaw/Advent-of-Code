package advent_of_code

import (
    "testing"
    "io/ioutil"
    "strings"
)

func TestDay4(t *testing.T) {

    bytes, _ := ioutil.ReadFile("day4_input.txt")
    inputText := string(bytes)

    // part 1
    count := NumberOfRealRooms(inputText)
    if count != 185371 {
        t.Errorf("Got %v expected %v", count, 185371)
    }

    // part 2
    decodedRooms := DecodeRealRooms(inputText)
    len := decodedRooms.Size()
    for i := 0; i < len; i++ {
    	decryptedRoom, _ := decodedRooms.Get(i)
    	s := decryptedRoom.(string)
    	if(strings.HasPrefix(s, "northpole")) {
    		split := strings.Split(s, " ")
    		sectorId := split[3]
		    if sectorId != "984" {
       			t.Errorf("Got %v expected %v", sectorId, "984")
    		}

    	}
    }
}
