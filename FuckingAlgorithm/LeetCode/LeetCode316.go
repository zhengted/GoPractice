package main

func removeDuplicateLetters(s string) string {
	nLength := len(s)
	lastIndex := make([]int, 26)
	for i := 0; i < nLength; i++ {
		lastIndex[s[i]-'a'] = i
	}
	var stack []uint8
	visited := make([]bool, 26)
	for i := 0; i < nLength; i++ {
		if visited[s[i]-'a'] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > s[i] &&
			lastIndex[stack[len(stack)-1]] > i {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i]-'a')
		visited[s[i]-'a'] = true
	}
	return string(stack)
}

func removeDuplicateLettersEx(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
