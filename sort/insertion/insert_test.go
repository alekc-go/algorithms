package insertion

import "testing"

var result []int

/*
pkg: algorithms/sort/insertion
BenchmarkInsert
BenchmarkInsert-8          	62818383	        18.8 ns/op
BenchmarkSortRecursive
BenchmarkSortRecursive-8   	22606267	        46.3 ns/op
PASS
*/
func BenchmarkInsert(b *testing.B) {
	testData := []int{3, 8, 5, 4, 1, 9, -2, 3, 8, 5, 4, 1, 9, -2}
	var r []int
	for n := 0; n < b.N; n++ {
		r = Sort(testData)
	}
	result = r
}
func BenchmarkSortRecursive(b *testing.B) {
	testData := []int{3, 8, 5, 4, 1, 9, -2, 3, 8, 5, 4, 1, 9, -2}
	var r []int
	for n := 0; n < b.N; n++ {
		r = SortRecursive(testData, 1, len(testData)-1)
	}
	result = r
}
