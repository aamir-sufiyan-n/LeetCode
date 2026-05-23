func check(nums []int) bool {
    count:=0
    for i:=0;i<len(nums)-1;i++{
        if nums[i]>nums[i+1]{
            count++
        }
    }
    if count ==1 && nums[len(nums)-1]<=nums[0]{return true}
    if count==0{ return true }
    return false
}