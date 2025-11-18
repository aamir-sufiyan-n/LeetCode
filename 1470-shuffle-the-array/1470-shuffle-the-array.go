func shuffle(nums []int, n int) []int {
    var res=make([]int,2*n)
    j:=0
    for i:=0;i<n;i++{
        res[j]=nums[i]
        res[j+1]=nums[n+i]
        j+=2
    }
    return res
}