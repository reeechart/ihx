package solution

import (
	"fmt"

	"github.com/reeechart/ihx/pkg/queue"
)

func MergeStrings(a string, b string) string {
	aq := queue.NewCharFromString(a)
	bq := queue.NewCharFromString(b)

	result := ""

	for !aq.IsEmpty() && !bq.IsEmpty() {
		charA, _ := aq.Dequeue()
		charB, _ := bq.Dequeue()

		result = fmt.Sprintf("%s%s%s", result, charA, charB)
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
