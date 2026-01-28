func sortColors(nums []int) {
    for i:= range nums{
        min:=i
        for j:=i+1;j<len(nums);j++{
            if nums[j]<nums[min]{min=j}
        }
        nums[min],nums[i]=nums[i],nums[min]
    }
    return 
}