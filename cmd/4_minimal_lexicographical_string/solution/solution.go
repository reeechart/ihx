package solution

import (
	"fmt"
	"strings"

	"github.com/reeechart/ihx/pkg/queue"
)

// GenerateLexicographicallyMinimalString generates a lexicographically minimum string based on two strings.
// It can only generate the string based on the principle of a queue (FIFO).
func GenerateLexicographicallyMinimalString(a string, b string) string {
	aq := queue.NewCharFromString(a)
	bq := queue.NewCharFromString(b)

	result := ""

	for !aq.IsEmpty() && !bq.IsEmpty() {
		charA, _ := aq.Peek()
		charB, _ := bq.Peek()

		switch strings.Compare(charA, charB) {
		case 1:
			result = fmt.Sprintf("%s%s", result, charB)
			bq.Dequeue()
		case -1:
			result = fmt.Sprintf("%s%s", result, charA)
			aq.Dequeue()
		default: // case 0
			// TODO: Decide which side to pick by peeking the next element, if still matches, pick any side
			result = fmt.Sprintf("%s%s", result, charA)
			aq.Dequeue()
		}
	}

	for !aq.IsEmpty() {
		char, _ := aq.Dequeue()
		result = fmt.Sprintf("%s%s", result, char)
	}

	for !bq.IsEmpty() {
		char, _ := bq.Dequeue()
		result = fmt.Sprintf("%s%s", result, char)
	}

	return result
}
