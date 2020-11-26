// Problem - https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	//abba
	stringMap := make(map[rune]int)
	start, longestSubstringLen := 0, 0
	for i, v := range s {
		if index, ok := stringMap[v]; ok {
			//Present
			if start < index+1 {
				start = index + 1
			}
		}
		stringMap[v] = i

		substringLen := i - start + 1
		if substringLen > longestSubstringLen {
			longestSubstringLen = substringLen
		}
	}
	return longestSubstringLen
}
