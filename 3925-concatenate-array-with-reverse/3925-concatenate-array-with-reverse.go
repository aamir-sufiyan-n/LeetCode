func concatWithReverse(nums []int) (res []int) {
    res=append(res,nums...)
    for j:=len(nums)-1;j>=0;j--{
        res=append(res,nums[j])
    }
    return 
}