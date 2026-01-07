package main

// longestZeroesSequence calculates longest sequence of zeroes in slice
func longestZeroes(seq []int) int {
	if len(seq) == 1 && seq[0] == 0 {
		return 1
	}
	seqLen := 0
	for i := 1; i < len(seq); i++ {
		if seq[i-1] == 0 {
			_seqLen := 1
			j := i
			for ; j < len(seq); j++ {
				if seq[j] == 0 {
					_seqLen++
				} else {
					break
				}
			}
			i = j
			if _seqLen > seqLen {
				seqLen = _seqLen
			}
		}
	}
	return seqLen
}
