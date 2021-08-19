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
			if pickNext(aq, bq) == 1 {
				result = fmt.Sprintf("%s%s", result, charA)
				aq.Dequeue()
			} else {
				result = fmt.Sprintf("%s%s", result, charB)
				bq.Dequeue()
			}
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

// PickNext returns either 1 or -1 by comparing two of the queues.
// Queue size needs to be guaranteed to be greater than 0, otherwise it will panic.
// Returns 1 when aq needs to be picked next.
// Returns -1 when bq needs to be picked next.
func pickNext(aq *queue.CharQueue, bq *queue.CharQueue) int {
	i := 0

	for i < aq.Size() && i < bq.Size() {
		a, _ := aq.Get(i)
		b, _ := bq.Get(i)

		switch strings.Compare(a, b) {
		case 1:
			return -1
		case -1:
			return 1
		default: // case 0
			i += 1
		}
	}

	for i < aq.Size() {
		a, _ := aq.Get(i)
		b, _ := bq.Get(0)

		switch strings.Compare(a, b) {
		case 1:
			return -1
		case -1:
			return 1
		default: // case 0
			i += 1
		}
	}

	for i < bq.Size() {
		a, _ := aq.Get(0)
		b, _ := bq.Get(i)

		switch strings.Compare(a, b) {
		case 1:
			return -1
		case -1:
			return 1
		default: // case 0
			i += 1
		}
	}

	// All characters are equal, return anything
	return 1
}
