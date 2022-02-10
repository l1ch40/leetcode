func search(nums []int, target int) int {
    for left, right := 0, len(nums) - 1; left <= right; {
        midIndex := left + (right - left) / 2
        if nums[midIndex] == target {
            return midIndex
        } else if nums[midIndex] < target {
            left = midIndex + 1
        } else {
            right = midIndex - 1
        }
    }
    return -1
}
