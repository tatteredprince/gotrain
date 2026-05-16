package main

func longestSubstrWoReapeats(str string) int {
	longestSubstr := 0
	chars := make(map[byte]struct{}, 0)
	for i := 0; i < len(str); i++ {
		if _, ok := chars[str[i]]; ok {
			if cnt := len(chars); cnt > longestSubstr {
				longestSubstr = cnt
			}
			chars = make(map[byte]struct{}, 0)
		}
		chars[str[i]] = struct{}{}
	}
	if cnt := len(chars); cnt > longestSubstr {
		return cnt
	}
	return longestSubstr
}
