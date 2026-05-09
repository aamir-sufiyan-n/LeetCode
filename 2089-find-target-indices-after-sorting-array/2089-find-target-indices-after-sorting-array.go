func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    var res []int
    for i,n:=range nums{
        if n==target{
            res=append(res,i)
        } 
    }
    return res
}