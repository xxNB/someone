func twoSum(nums []int, target int) []int {
    maps := make(map[int]int)
    for i, _ := range (nums){
        complement := target - nums[i]
        
        if _, ok := maps[complement]; ok {
            return []int{maps[complement], i}
        }
        maps[nums[i]]=i
    }
    return nil
}