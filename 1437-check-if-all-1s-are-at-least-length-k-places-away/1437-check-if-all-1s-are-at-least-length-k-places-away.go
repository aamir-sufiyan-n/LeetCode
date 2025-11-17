func kLengthApart(nums []int, k int) bool {
    prev:=-1
    for i:=0;i<len(nums);i++{
        if prev==-1 && nums[i]==1{
            prev=i
        }else if nums[i]==1 && i-prev<=k{
            return false
        }
    }
    return true
}
