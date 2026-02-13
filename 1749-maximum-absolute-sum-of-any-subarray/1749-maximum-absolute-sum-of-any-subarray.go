func maxAbsoluteSum(nums []int) int {
    maxSum:=0
    minSum:=0
    curr:=0
    for _,n:=range nums{
        curr+=n
        maxSum=max(maxSum,curr)
        minSum=min(curr,minSum)
    }
    return maxSum-minSum
}
func max(a,b int )int{
    if a>b{ return a}
    return b
}
func min(a,b int)int{
    if a>b{ return b}
    return a
}