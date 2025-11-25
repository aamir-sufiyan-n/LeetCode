func findMissingElements(nums []int) []int {
    m:=make(map[int]bool)
    min,max:=nums[0],nums[0]
    res:=[]int{}
    for _,n:=range nums{
        m[n] = true
        if n<min { min=n }
        if n>max { max=n }
    }
   for i:=min;i<=max;i++{
        if !m[i] { res=append(res,i) }
    }
return res
}