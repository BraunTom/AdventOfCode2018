package SixteenthDay

func target(instructions, after [4]int) int {
	return after[instructions[3]]
}

func registerA(instruction, before [4]int) int {
	return before[instruction[1]]
}

func registerB(instruction, before [4]int) int {
	return before[instruction[2]]
}

func valueA(instruction [4]int) int {
	return instruction[1]
}

func valueB(instruction [4]int) int {
	return instruction[2]
}

func untouched(before, instruction, after [4]int) bool {
	for i := 0; i < 4; i++ {
		if i != instruction[3] && before[i] != after[i] {
			return false
		}
	}

	return true
}

func testAddr(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)+registerB(instruction, before)
}

func testAddi(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)+valueB(instruction)
}

func testMulr(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)*registerB(instruction, before)
}

func testMuli(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)*valueB(instruction)
}

func testBanr(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)&registerB(instruction, before)
}

func testBani(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)&valueB(instruction)
}

func testBorr(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)|registerB(instruction, before)
}

func testBori(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)|valueB(instruction)
}

func testSetr(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == registerA(instruction, before)
}

func testSeti(before, instruction, after [4]int) bool {
	return untouched(before, instruction, after) && target(instruction, after) == valueA(instruction)
}

func testGtir(before, instruction, after [4]int) bool {
	if valueA(instruction) > registerB(instruction, before) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

func testGtri(before, instruction, after [4]int) bool {
	if registerA(instruction, before) > valueB(instruction) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

func testGtrr(before, instruction, after [4]int) bool {
	if registerA(instruction, before) > registerB(instruction, before) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

func testEqir(before, instruction, after [4]int) bool {
	if valueA(instruction) == registerB(instruction, before) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

func testEqri(before, instruction, after [4]int) bool {
	if registerA(instruction, before) == valueB(instruction) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

func testEqrr(before, instruction, after [4]int) bool {
	if registerA(instruction, before) == registerB(instruction, before) {
		return untouched(before, instruction, after) && target(instruction, after) == 1
	} else {
		return untouched(before, instruction, after) && target(instruction, after) == 0
	}
}

var tests = []func([4]int, [4]int, [4]int) bool{
	testAddr, testAddi,
	testMulr, testMuli,
	testBanr, testBani,
	testBorr, testBori,
	testSetr, testSeti,
	testGtir, testGtri, testGtrr,
	testEqir, testEqri, testEqrr}
