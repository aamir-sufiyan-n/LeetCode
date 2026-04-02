func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    count:=0
    for _,n:=range nums1{
        for _,n2:=range nums2{
            if n%(n2*k)==0{
                count++
            }
        }
    }
    return count  
}