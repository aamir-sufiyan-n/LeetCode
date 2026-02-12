func containsNearbyDuplicate(nums []int, k int) bool {
    checker:=make(map[int]int)
    for i,n:=range nums{
        if index,ok:=checker[n];ok{
            if abs(index-i)<=k{return true}
        }
        checker[n]=i
    }
    return false
}
func abs(x int)int{
    if x<0{return x*-1}
    return x
}