func findDuplicates(nums []int) []int {
    var res[]int
    // count:=make(map[int]int)
    // for _,num:=range nums {
    //     count[num]++
    //     if val,found:= count[num]; found && val>1 {
    //         res=append(res,num)
    //         }
    // }
    for _,num:=range nums{
        index:=abs(num)-1
        if nums[index]<0{
            res=append(res,abs(num))
        }else{
            nums[index]=nums[index]*(-1)
        }
    }
    return res
}
func abs(x int)int{
    if x<0 {
    return -x }
    return x
}