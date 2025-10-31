func getSneakyNumbers(nums []int) []int {
   result:=[]int{}
   for i:=0;i<len(nums);i++{
    for j:=i+1;j<len(nums);j++{
        if nums[i]==nums[j]{
            result=append(result,nums[i])
        }
    }
   }
   return result
}