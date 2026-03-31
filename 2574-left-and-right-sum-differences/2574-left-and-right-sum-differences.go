func leftRightDifference(nums []int) []int {
    leftSum:=0
    rightSum:=Sum(nums)
    
    var total = []int{}
    for _,n:=range nums{
        rightSum-=n
        total=append(total,abs(rightSum-leftSum))
        leftSum+=n
    }
    return total
}
func abs(n int)int{
    if n<0{ return n *-1 }
    return n
}
func Sum(nums[]int)int{
    var sum int
    for _,v:=range nums{
        sum+=v
    }
    return sum
}