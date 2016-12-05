package advent_of_code


func move1(direction byte, pos *Point) {
	switch direction {
		case 'U':
			if(pos.y > 0) {
				pos.y--;
			}
		case 'R':
			if(pos.x < 2) {
				pos.x++
			}
		case 'L':
			if(pos.x > 0) {
				pos.x--
			}
		case 'D':
			if(pos.y < 2) {
				pos.y++
			}
	}
}

// 1 2 3
// 4 5 6
// 7 8 9
func ButtonMap1() map[Point]string {
	m := make(map[Point]string)
	m[Point{x : 0, y : 0}] = "1"
	m[Point{x : 0, y : 1}] = "4"
	m[Point{x : 0, y : 2}] = "7"

	m[Point{x : 1, y : 0}] = "2"
	m[Point{x : 1, y : 1}] = "5"
	m[Point{x : 1, y : 2}] = "8"

	m[Point{x : 2, y : 0}] = "3"
	m[Point{x : 2, y : 1}] = "6"
	m[Point{x : 2, y : 2}] = "9"

	return m
}


func move2(direction byte, pos *Point) {
	switch direction {
		case 'U':
			if(((pos.x == 1 || pos.x == 3) && pos.y > 1) || (pos.x == 2 && pos.y > 0)) {
				pos.y--
			}
		case 'R':
			if(((pos.y == 1 || pos.y == 3) && pos.x < 3) || (pos.y == 2 && pos.x < 4)) {
				pos.x++
			}
		case 'L':
			if(((pos.y == 1 || pos.y == 3) && pos.x > 1) || (pos.y == 2 && pos.x > 0)) {
				pos.x--
			}
		case 'D':
			if(((pos.x == 1 || pos.x == 3) && pos.y < 3) || (pos.x == 2 && pos.y < 4)) {
				pos.y++
			}
	}
}

//     1
//   2 3 4
// 5 6 7 8 9
//   A B C
//     D
func ButtonMap2() map[Point]string {
	m := make(map[Point]string)
	m[Point{x : 0, y : 2}] = "5"

	m[Point{x : 1, y : 1}] = "2"
	m[Point{x : 1, y : 2}] = "6"
	m[Point{x : 1, y : 3}] = "A"

	m[Point{x : 2, y : 0}] = "1"
	m[Point{x : 2, y : 1}] = "3"
	m[Point{x : 2, y : 2}] = "7"
	m[Point{x : 2, y : 3}] = "B"
	m[Point{x : 2, y : 4}] = "D"

	m[Point{x : 3, y : 1}] = "4"
	m[Point{x : 3, y : 2}] = "8"
	m[Point{x : 3, y : 3}] = "C"

	m[Point{x : 4, y : 2}] = "9"

	return m
}

func ButtonMap(partOne bool) map[Point]string {
	if(partOne) {
		return ButtonMap1()
	} else {
		return ButtonMap2()
	}
}

func startingPosition(partOne bool) *Point {
	if(partOne) {
		return &Point{x : 1, y : 1}
	} else {
		return &Point{x : 0, y : 2}
	}
}


func executeCommand(command string, pos *Point, partOne bool) {
	str_len := len(command)
	for i := 0; i < str_len; i++ {
		if(partOne) {
			move1(command[i], pos)
		} else {
			move2(command[i], pos)
		}
	}
}




func findBathroom(commands []string, partOne bool) []string {
	m := ButtonMap(partOne)
	position := startingPosition(partOne)

	code := make([]string, len(commands))
	for index, command := range commands {

		executeCommand(command, position, partOne)
		code[index] = m[*position]
	}
	return code
}
