package solution

import (
	"github.com/reeechart/ihx/cmd/5_list_max/model"
)

// MaximumOfList returns a 64-bit integer that reprensents the maximum of value in a given array
// after doing some operations given.
func MaximumOfList(elemCount int, operations []model.Operation) int64 {
	arr := NewZeroArray(elemCount)

	for _, op := range operations {
		arr.Operate(op)
	}

	return arr.Max()
}

type Array struct {
	arr []int64
}

func NewZeroArray(length int) Array {
	arr := make([]int64, length)
	return Array{arr}
}

func (a *Array) Operate(operation model.Operation) {
	for i := operation.LeftBound - 1; i < operation.RightBound; i++ {
		a.arr[i] += operation.Addition
	}
}

func (a *Array) Max() int64 {
	max := int64(0)

	for _, val := range a.arr {
		if max < val {
			max = val
		}
	}

	return max
}
