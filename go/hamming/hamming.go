package hamming

import "errors"

const testVersion = 4

//Distance returns the Hamming distance between two nucleotide strings
func Distance(a, b string) (int, error) {
	var count int

	if len(a) != len(b) {
		return count, errors.New("Nucleotide strings must be the same length")
	}

	for index, base := range a {
		if byte(base) != b[index] {
			count++
		}
	}
	return count, nil
}