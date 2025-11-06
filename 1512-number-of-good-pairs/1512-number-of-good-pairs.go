func numIdenticalPairs(nums []int) int {
    some:=make(map[int]int)
    for _,num:=range nums{
        some[num]++
    }
    res:=0
    for _,freq:=range some{
        if freq>1{
            res+=freq*(freq-1)/2
        }
    }
    return res
}