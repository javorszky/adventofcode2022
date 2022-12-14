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

type item interface {
	String() string
	Day13()
}

type integer int

func (i integer) String() string {
	return fmt.Sprintf(" %d ", i)
}

func (i integer) Day13() {
	// do nothing
}

type list []item

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
