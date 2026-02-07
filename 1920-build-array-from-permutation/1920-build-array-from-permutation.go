func buildArray(nums []int) []int {
    var res=[]int{}
    for i:= range nums{
        res=append(res,nums[nums[i]])
    }
    return res
}