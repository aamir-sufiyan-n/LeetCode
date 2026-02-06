func minimumCost(nums []int) int {
    min1,min2:=math.MaxInt,math.MaxInt
    for i:=1;i<len(nums);i++{
        if nums[i]<min1 {
            min1,min2=nums[i],min1
        }else if nums[i]<min2{
            min2=nums[i]
        }
    }
    return nums[0] + min1 + min2
}