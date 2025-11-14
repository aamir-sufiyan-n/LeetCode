func findDisappearedNumbers(nums []int) []int {
    res:=[]int{}
    for _,num:=range nums{
        index:=abs(num)-1
        if nums[index]>0{nums[index]*=-1}
    }
    for i,n:=range nums{
        if n>0{res=append(res,i+1)}
    }
    return res
}
func abs(x int)int{
    if x<0 {return x*-1}
    return x
}

