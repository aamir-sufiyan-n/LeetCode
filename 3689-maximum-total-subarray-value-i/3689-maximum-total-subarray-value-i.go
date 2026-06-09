func maxTotalValue(nums []int, k int) int64 {
    max:=nums[0]
    min:=nums[0]
    for _,n:=range nums{
        if n>max { max=n }
        if n<min { min=n }
    }
    return int64(k*(max-min))
}