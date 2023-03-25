package strcmp

import (
	"fmt"
	"strings"
)

type DiffType int

const (
	Addition    DiffType = 1
	Subtraction DiffType = 2
)

const (
	Red   = 32
	Green = 31
)

type Diff struct {
	Text string
	Line int
	Type DiffType
}

func ANSIDiff(new string, old string) string {
	xs := strings.Split(old, "\n")
	ys := strings.Split(new, "\n")

	dpTable := LCS(xs, ys)

	return BuildDiff(dpTable, xs, ys, true)
}

// LCS find the longest common substr with Dynamic Programing.
// Git diff is the reverse of LCS.
// more: https://riddhinn.medium.com/create-a-text-differentiator-using-longest-common-subsequence-afc431b105dd
func LCS(x []string, y []string) [][]int {
	var dpTable [][]int

	for i := 0; i <= len(x); i++ {
		dpTable = append(dpTable, make([]int, len(y)+1))
	}

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				dpTable[i][j] = dpTable[i-1][j-1] + 1
			} else {
				dpTable[i][j] = max(dpTable[i-1][j], dpTable[i][j-1])
			}
		}
	}

	return dpTable
}

// BuildDiff traceback the LCS dp table and check the diff two input string.
func BuildDiff(dpTable [][]int, x []string, y []string, pretty bool) string {
	var diffReverse []Diff

	i := len(x)
	j := len(y)

	for i > 0 && j > 0 {
		if dpTable[i][j] == dpTable[i-1][j] {
			diffReverse = append(diffReverse, Diff{
				Text: x[i-1],
				Line: i,
				Type: Addition,
			})

			i--
		} else if dpTable[i][j] == dpTable[i][j-1] {
			diffReverse = append(diffReverse, Diff{
				Text: y[j-1],
				Line: j,
				Type: Subtraction,
			})

			j--
		} else {
			i--
			j--
		}
	}

	for i > 0 {
		diffReverse = append(diffReverse, Diff{
			Text: x[i-1],
			Line: i,
			Type: Addition,
		})

		i -= 1
	}
	for j > 0 {
		diffReverse = append(diffReverse, Diff{
			Text: y[j-1],
			Line: j,
			Type: Subtraction,
		})

		j -= 1
	}

	for i, j := 0, len(diffReverse)-1; i < j; i, j = i+1, j-1 {
		diffReverse[i], diffReverse[j] = diffReverse[j], diffReverse[i]
	}

	return diffJoin(diffReverse, pretty)
}

func diffJoin(diffs []Diff, pretty bool) string {
	var strBuff strings.Builder

	for _, diff := range diffs {
		switch diff.Type {
		case Addition:
			if pretty {
				strBuff.WriteString(fmt.Sprintf("%s %d+ %s%s\n", escapeCode(Red), diff.Line, diff.Text, escapeCode(0)))
			} else {
				strBuff.WriteString(diff.Text)
			}

		case Subtraction:
			if pretty {
				strBuff.WriteString(fmt.Sprintf("%s %d- %s%s\n", escapeCode(Green), diff.Line, diff.Text, escapeCode(0)))
			} else {
				strBuff.WriteString(diff.Text)
			}
		}

	}

	return strBuff.String()
}

func escapeCode(code int) string {
	return fmt.Sprintf("\x1b[%dm", code)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
