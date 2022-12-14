package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// [ - 91
// ] - 93
// , - 44
// 0 - 48
// 1 - 49
// 2 - 50
// 3 - 51
// 4 - 52
// 5 - 53
// 6 - 54
// 7 - 55
// 8 - 56
// 9 - 57
const (
	charLeftSquare  uint8 = 91
	charRightSquare uint8 = 93
	charComma       uint8 = 44

	correctOrder = iota
	incorrectOrder
	continueEvaluation

	typeList
	typeInteger
)

var (
	errNotANumber = errors.New("character is not a number")
)

func parseLine(line string, d int) (list, int, error) {
	l := list{}
	var sb strings.Builder
	// fmt.Printf("\n%sparsing line '%s' len(%d), skipping first char\n", strings.Repeat(" ", d*2), line, len(line))

	for i := 1; i <= len(line); i++ {
		// fmt.Printf("%s%d - %s - %d\n", strings.Repeat(" ", d*2), i, string(line[i]), line[i])

		switch line[i] {
		case charLeftSquare:
			rest := line[i:]
			// fmt.Printf("%sleft square, passing the rest of the line to parseline: '%s'\n", strings.Repeat(" ", d*2), rest)

			sp, j, err := parseLine(rest, d+1)
			if err != nil {
				return list{}, 0, err
			}
			// fmt.Printf("%sappending '%s' (%#v) to current list, previous index was %d, upped by %d -1 , new index is %d\n", strings.Repeat(" ", d*2), sp, sp, i, j, i+j-1)

			i = i + j - 1
			l = append(l, sp)
		case charRightSquare:
			strToParse := sb.String()
			if strToParse != "" {
				sp, err := strconv.Atoi(sb.String())
				if err != nil {
					return list{}, 0, errors.Wrapf(err, "could not parse collected number %s into int", sb.String())
				}
				// fmt.Printf("%sright square, collecting the number '%d', adding and returning\n", strings.Repeat(" ", d*2), sp)

				l = append(l, integer(sp))
				sb.Reset()
			} else {
				// fmt.Printf("%sright square, returning without collecting anything\n", strings.Repeat(" ", d*2))
			}

			return l, i + 1, nil
		case charComma:
			strToParse := sb.String()
			if strToParse != "" {
				sp, err := strconv.Atoi(sb.String())
				if err != nil {
					return list{}, 0, errors.Wrapf(err, "could not parse collected number %s into int", sb.String())
				}
				// fmt.Printf("%scomma, appending the number '%d', adding and resetting\n", strings.Repeat(" ", d*2), sp)

				l = append(l, integer(sp))
				sb.Reset()
			}
		default:
			if line[i] < 48 || line[i] > 57 {
				return list{}, 0, errNotANumber
			}
			sb.WriteString(string(line[i]))
		}
	}
	// fmt.Printf("%sreturning from main parseLine", strings.Repeat(" ", d*2))
	return l, len(line), nil
}

// smallerList compares two lists, and returns which one is the shorter. It's using
// three possible return values as iota consts:
//
// - correctOrder
// - incorrectOrder
// - continueEvaluation
func smallerList(left, right list) int {
	fallback := continueEvaluation

	le := len(left)
	lr := len(right)
	min := le

	if le < lr {
		fallback = correctOrder
	} else if lr < le {
		min = lr
		fallback = incorrectOrder
	}

	for i := 0; i < min; i++ {
		le := left[i]
		re := right[i]
		switch le.Type() {
		case typeInteger:
			switch re.Type() {
			case typeInteger:
				// both integers, if left integer is lower, correct order
				// right integer is lower, inputs are not in the same order
				// integers same --> continue
				if le.(integer) == re.(integer) {
					// they are equal, no decision
					continue
				}

				if le.(integer) < re.(integer) {
					return correctOrder
				}

				return incorrectOrder
			case typeList:
				// left is integer, right is list
				// if comparing list vs integer, make the integer into a list
				// then compare lists as above
				le = list{le}
				result := smallerList(le.(list), re.(list))
				if result == continueEvaluation {
					continue
				}

				return result
			}
		case typeList:
			switch re.Type() {
			case typeList:
				// both are lists, enter list, and start comparing values
				// if left runs out of items first, correct order
				// if right runs out of items first, incorrect order
				// same length, contents also don't decide --> continue
				result := smallerList(le.(list), re.(list))
				if result == continueEvaluation {
					continue
				}

				return result
			case typeInteger:
				// right is integer, left is list
				// if comparing list vs integer, make the integer into a list
				// then compare lists as above
				re = list{re}
				result := smallerList(le.(list), re.(list))
				if result == continueEvaluation {
					continue
				}

				return result
			}
		}
	}

	return fallback
}

type item interface {
	String() string
	Day13()
	Type() int
}

type integer int

func (i integer) Type() int {
	return typeInteger
}

func (i integer) String() string {
	return fmt.Sprintf(" %d ", i)
}

func (i integer) Day13() {
	// do nothing
}

type list []item

func (l list) Type() int {
	return typeList
}

func (l list) String() string {
	s := make([]string, len(l))
	for i, n := range l {
		s[i] = n.String()
	}
	return fmt.Sprintf("[ %s ]", strings.Join(s, ", "))
}

func (l list) Day13() {
	// do nothing
}
