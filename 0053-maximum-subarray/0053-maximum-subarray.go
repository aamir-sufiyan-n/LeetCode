func maxSubArray(nums []int) int {
maxSum:=nums[0]
currentSum:=nums[0]
for i:=1;i<len(nums);i++{
    // if nums[i]>currentSum+nums[i]{
    //     currentSum=nums[i]
    // }else{currentSum+=nums[i]}

    // if currentSum>maxSum { maxSum=currentSum  }
    
    currentSum=max(nums[i],currentSum+nums[i])
    maxSum=max(currentSum,maxSum)
}
return maxSum
}