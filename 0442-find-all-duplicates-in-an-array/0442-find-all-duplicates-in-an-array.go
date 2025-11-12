func findDuplicates(nums []int) []int {
    var res[]int
    count:=make(map[int]int)
    for _,num:=range nums {
        count[num]++
        if val,found:= count[num]; found && val==2 {res=append(res,num)}
    }
    
    return res
}