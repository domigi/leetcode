func twoSum(nums []int, target int) []int {
    diffs := map[int]int{}
    for i := 0; i < len(nums); i++ {
        current := nums[i]
                if index, exists := diffs[current]; exists{
            return []int{index, i}
        }
        dif := target - current
        diffs[dif] = i
    }
    return nil
}