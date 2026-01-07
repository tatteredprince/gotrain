package main

// monotonicSequence indicates that sequence passed monotonically increases or descreases
func monotonicSequence(seq []int) bool {
	if len(seq) < 2 {
		return false
	}
	idx := 1
	isAscending := false
	for ; idx < len(seq); idx++ {
		if seq[idx-1] != seq[idx] {
			if seq[idx-1] < seq[idx] {
				isAscending = true
			}
			break
		}
	}
	for ; idx < len(seq); idx++ {
		if isAscending && seq[idx-1] > seq[idx] || !isAscending && seq[idx-1] < seq[idx] {
			return false
		}
	}
	return true
}
