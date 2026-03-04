func sumOfUnique(nums []int) int {
    checker:=make(map[int]int)
    sum:=0
    for _,n:=range nums{
        checker[n]++
    }
    for k,v:=range checker{
        if v ==1{sum+=k}
    }
    return sum
}