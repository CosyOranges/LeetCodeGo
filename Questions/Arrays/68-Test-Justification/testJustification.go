package main

import (
	"fmt"
)

func FullJustify(words []string, maxwidth int) []string {
	ans := []string{}
	start, stop := 0, 0
	length := 0
	line := ""

	for i := 0; i < len(words); i++ {
		if length+len(" "+words[i]) > maxwidth && i > 0 {
			stop = i
			spacesLeft := maxwidth - length
			line = words[start]
			for j := start + 1; j < stop; j++ {
				line += " " + words[j]
			}
			// fmt.Printf("Line: %v | Number of words: %v | spaces left: %v\n", line, stop-start, spacesLeft)
			// temp := lineBuilder(spacesLeft, line)
			// fmt.Printf("Line Builder: %v| Length of that: %v\n", temp, len(temp))
			if spacesLeft > 0 && stop-start > 1 {
				ans = append(ans, lineBuilder(spacesLeft, line, false))
			} else if spacesLeft > 0 && stop-start == 1 {
				ans = append(ans, lineBuilder(spacesLeft-1, line+" ", false))
			} else {
				ans = append(ans, line)
			}

			start = i
			length = len(words[i])
		} else if length == 0 {
			length += len(words[i])
		} else {
			length += len(" " + words[i])
		}

		if i == len(words)-1 {
			stop = i
			spacesLeft := maxwidth - length
			line = words[start]
			for j := start + 1; j < stop+1; j++ {
				line += " " + words[j]
			}

			if spacesLeft > 0 {
				ans = append(ans, lineBuilder(spacesLeft-1, line+" ", true))
			} else {
				ans = append(ans, line)
			}

			// fmt.Printf("Line: %v | Number of words: %v | spaces left: %v\n", line, stop-start, spacesLeft)
			// fmt.Printf("Line Builder: %v| Length of that: %v\n", temp, len(temp))
		}
	}

	return ans
}

func lineBuilder(spaces int, line string, lastline bool) string {
	if !lastline {
		for spaces > 0 {
			needToAddSpace := true
			for i := 0; i < len(line); i++ {
				if line[i] == ' ' && needToAddSpace {
					line = line[:i] + " " + line[i:]
					spaces--
					needToAddSpace = false
				} else if line[i] != ' ' {
					needToAddSpace = true
				}
				if spaces == 0 {
					break
				}
			}
		}
	} else {
		for spaces > 0 {
			line += " "
			spaces--
		}
	}

	return line
}

func main() {
	test := []string{"Listen", "to", "many,", "speak", "to", "a", "few."}
	just := FullJustify(test, 6)

	for _, x := range just {
		fmt.Printf("%v | Length: %v\n", x, len(x))
	}
	// fmt.Printf("Justified:\n%v\n%v\n%v\n", just[0], just[1], just[2])
}
