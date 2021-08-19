package solution

// Swap function swaps two integers in-place.
func Swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}
