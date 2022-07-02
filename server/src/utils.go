package src

//URemoveItem removes item from slice and returns the slice in no particular order. This is a fast way to remove an item
func URemoveItem[T any](s []T, k int) []T {
	s[k] = s[len(s)-1] // Copy last element to index k.
	s = s[:len(s)-1]   // Truncate slice.
	return s
}

//URemoveItems removes items from slice and returns the slice in no particular order. This is a fast way to remove an item
func URemoveItems[T any](s []T, k []int) []T {
	for _, v := range k {
		s[v] = s[len(s)-len(k)] // Copy last element to index k. // Truncate slice.
	}
	s = s[:len(s)-len(k)]
	return s
}
