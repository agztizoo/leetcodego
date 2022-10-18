package easy

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        numMap[num] = i
    }
    for i, num := range nums {
        if num > target {
            continue
        }
        v := target - num
        if idx, exists := numMap[v]; exists && idx != i {
            return []int{i, idx}
        }
    }
    return nil
}
