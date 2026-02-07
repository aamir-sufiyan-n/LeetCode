func numIdenticalPairs(nums []int) int {
    // some:=make(map[int]int)
    // for _,num:=range nums{
    //     some[num]++
    // }
    res:=0
    // for _,freq:=range some{
    //     if freq>1{
    //         res+=freq*(freq-1)/2
    //     }
    // }

    for i:=0;i<len(nums);i++{
        for j:=0;j<len(nums);j++{
            if nums[i]==nums[j] && i<j{
                res++
            }
        }
    }
    return res
}