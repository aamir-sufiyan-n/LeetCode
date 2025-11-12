func missingNumber(nums []int) int {
    n:=len(nums)
    trueSum:=(n*(n+1))/2
    var sum int
    for _,num:= range nums{sum+=num}
    return trueSum-sum
}