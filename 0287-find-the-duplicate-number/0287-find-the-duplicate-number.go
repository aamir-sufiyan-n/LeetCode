func findDuplicate(nums []int) int {
    checker:=make(map[int]bool)
    for _,n:=range nums{
        if checker[n]{return n}
        checker[n]=true
    }
    return -1

    //will learn floyd's later
}