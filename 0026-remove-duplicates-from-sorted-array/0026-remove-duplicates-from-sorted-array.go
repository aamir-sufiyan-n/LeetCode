func removeDuplicates(nums []int) int {
    i,j:=0,1
    l:=len(nums)
    for j<l{
        if nums[i]!=nums[j]{
            i++
            nums[i]=nums[j]
        }

        j++
    }
    return i+1
}