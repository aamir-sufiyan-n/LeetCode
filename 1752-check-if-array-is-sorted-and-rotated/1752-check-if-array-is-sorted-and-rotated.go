func check(nums []int) bool {
    count:=0
    for i:=0;i<len(nums)-1;i++{
        if nums[i]>nums[i+1]{
            count++
        }
    }
    if nums[0]<nums[len(nums)-1]{
        count++
    }
    return count<=1
}