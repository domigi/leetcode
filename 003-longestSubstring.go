func lengthOfLongestSubstring(s string) int {
    start := 0
    chars := map[rune]int{}
    longest := 0
    current := 0
    for index, char := range s {
        if previndex, contains := chars[char]; previndex >= start && contains {
            start = previndex + 1
        }
        current = index + 1 - start
        if current > longest {
            longest = current
        }
        chars[char] = index
    }
    return longest
}