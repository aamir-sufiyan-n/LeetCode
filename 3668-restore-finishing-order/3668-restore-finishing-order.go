func recoverOrder(order []int, friends []int) []int {
    count:=make(map[int]bool)
    res:=[]int{}
    for _,p:=range friends{
        count[p]=true
    }
    for _,n:=range order{
        if count[n]{res=append(res,n)}
    }
    return res
}