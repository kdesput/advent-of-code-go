package main

func gravityAssistProgram(inputCode []int) int {
	var currPos int

	for currPos >= 0 {
		switch inputCode[currPos] {
		case 1:
			inputCode[inputCode[currPos+3]] = inputCode[inputCode[currPos+1]] + inputCode[inputCode[currPos+2]]
			currPos += 4
		case 2:
			inputCode[inputCode[currPos+3]] = inputCode[inputCode[currPos+1]] * inputCode[inputCode[currPos+2]]
			currPos += 4
		case 99:
			currPos = -1
		}
	}
	return inputCode[0]
}
