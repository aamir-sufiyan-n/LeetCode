func separateDigits(nums []int) []int {
    var res []int
    for _,n:=range nums{
        var digits []int
        for n>0{
            digits=append([]int{n%10},digits...)
            n/=10
        }
        res=append(res,digits...)
    }
    return res
}