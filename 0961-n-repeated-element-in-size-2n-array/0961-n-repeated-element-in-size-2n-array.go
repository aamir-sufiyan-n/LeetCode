func repeatedNTimes(nums []int) int {
    check:=make(map[int]int)
    for _,n:=range nums{
        check[n]++
        if check[n]>1{
            return n
        }
    }
    return -1
}