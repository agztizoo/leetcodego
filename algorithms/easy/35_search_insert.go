package easy

// 35. Search Insert Position
func SearchInsert(nums []int, target int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    begin, end := 0, len(nums)
    if nums[begin] > target {
        return begin
    }
    if nums[end-1] < target {
        return end
    }
    for begin <= end {
        mid := (begin + end) / 2
        if nums[mid] == target {
            return mid
        }
        if target < nums[mid] {
            if mid > 0 && target > nums[mid-1] {
                return mid
            } else {
                end = mid - 1
            }
        } else {
            if mid < len(nums)-1 && target < nums[mid+1] {
                return mid + 1
            } else {
                begin = mid + 1
            }
        }
    }
    return 0
}
