func intersection(nums1 []int, nums2 []int) []int {
    look:=make(map[int]bool)
    res:=make([]int,0)
    for _,n:=range nums1{
        look[n]=true
    }
    for _,num:= range nums2{
        if look[num] && !slices.Contains(res,num){
            res=append(res,num)
        }
    }
    return res
}