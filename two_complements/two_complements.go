package main

// hasTwoComplements indicates that slice contains two values which in sum gives target value
func hasTwoComplements(target int, seq []int) bool {
	set := map[int]struct{}{}
	for _, v := range seq {
		compl := target - v
		if _, ok := set[compl]; ok == true {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
