package util

// Splice removes the specified element from the slice by its index.
// If the provided index is outside the boundaries of the slice, it proceeds to return the original slice.
// e.g:
// Splice([]int{9, 8, 7, 6, 5}, 2) = []int{9, 8, 6, 5}
// Splice([]float32{1.44, 9.99, 3.14}, 0) = []float32{9.99, 3.14}
// Splice([]string{"apples", "oranges", "banana", "mango"}, 4) = []string{"apples", "oranges", "banana", "mango"}
// Splice([]byte{0, 255}, -1) = []byte{0, 255}


func Splice[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {   // If the index is outside the boundaries of the slice, it proceeds to return the original slice.
		return slice
	}
	return append(slice[:index], slice[index+1:]...)   // Append the elements of the slice after the index to the slice before the index.
}