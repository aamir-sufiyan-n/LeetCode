func findDuplicate(nums []int) int {
    // checker:=make(map[int]bool)
    // for _,n:=range nums{
    //     if checker[n]{return n}
    //     checker[n]=true
    // }
    // return -1

     i:=0
     for i<len(nums){
        if nums[i]==i{continue}
        if nums[i]==nums[nums[i]]{ return nums[i]}
        nums[i],nums[nums[i]]=nums[nums[i]],nums[i]
     }
     return -1
    //will learn floyd's later
}