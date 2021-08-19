package model

// Operation represents the a (left bound), b (right bound), k (addition) in the input stream
type Operation struct {
	LeftBound  int
	RightBound int
	Addition   int64
}
