package main

// TODO:I do not fix this question
func predictPartyVictory(senate string) string {
	var radient, dire []int
	for i, b := range senate {
		if b == 'R' {
			radient = append(radient, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radient) > 0 && len(dire) > 0 {
		if radient[0] < dire[0] {
			radient = append(radient, radient[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radient = radient[1:]
		dire = dire[1:]
	}
	if len(radient) > 0 {
		return "Radient"
	}
	return "Dire"
}
