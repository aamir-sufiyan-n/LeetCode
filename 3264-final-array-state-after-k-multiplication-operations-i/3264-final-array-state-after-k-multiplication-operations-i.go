func getFinalState(nums []int, k int, multiplier int) []int {
    for k !=0{
        minIndex:=0
        for i:=0;i<len(nums);i++{
            if nums[i]<nums[minIndex]{
                minIndex=i
            }
        }
        nums[minIndex]*=multiplier
        k--
    }
    return nums
}