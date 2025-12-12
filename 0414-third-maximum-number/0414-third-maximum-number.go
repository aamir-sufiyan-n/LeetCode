func thirdMax(nums []int) int {
    check:=make(map[int]bool)
    for _,n:=range nums{
        check[n]=true
    }
    res:=[]int{}
    for n:=range check{
        res=append(res,n)
    }
    sort.Ints(res)
    if len(res)>=3{
        return res[len(res)-3]
    }
    return res[len(res)-1]
}