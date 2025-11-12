func findDuplicates(nums []int) []int {
    var res[]int
    count:=make(map[int]int)
    for _,num:=range nums {
        count[num]++
    }
    for n,c:= range count{
        if c>1{res=append(res,n)}
    }
    return res
}