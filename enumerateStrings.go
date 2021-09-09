package main

import "fmt"

func main() {
	fmt.Println(enumerateStrings([]rune{'a', 'b'}, 2))
}

func enumerateStrings(charSet []rune, maxLength int) []string {
	output := make([]string, 0)
	for i := 1; i <= maxLength; i++ {
		output = append(output, allStringsOfLength(i, charSet)...)
	}
	return output
}

func allStringsOfLength(length int, charSet []rune) []string {
	output := make([]string, 0)

	maxChar := len(charSet) - 1
	allMax := make([]int, length)
	for i := range allMax {
		allMax[i] = maxChar
	}

	idx := length - 1
	indexList := make([]int, length)

	for !areEqual(indexList,allMax) {
		p := makePassword(indexList, charSet)
		idx, indexList = next(idx, indexList, maxChar)
		output = append(output, p)
	}

    // Add final password
    output = append(output, makePassword(indexList, charSet))
   
	return output
}

func next(idx int, indexList []int, maxChar int) (int, []int) {
	// Standard case
	if indexList[idx] < maxChar {
		indexList[idx]++
		return idx, indexList
	}

	// If maxChar reached
	if idx > 0 {
		idx, indexList = next(idx-1, indexList, maxChar)
		idx++
		return idx, indexList
	}

	// If on the first character, increment the from the last character
	idx, indexList = next(len(indexList)-1, indexList, maxChar)
	idx = 0
	return idx, indexList
}

func makePassword(indexList []int, charSet []rune) string {
    p := ""
    for _, index := range indexList {
        p += string(charSet[index])
    }    
    return p
}

func areEqual(first, second []int) bool {
    if len(first) != len(second) {
        return false
    }
    for i, r := range first {
        if r != second[i] {
            return false
        }
    } 
    return true
}
