func twoSum(nums []int, target int) []int {
   m:=make(map[int]int)
   for i,n:=range nums{
     if j,exist:=m[target-n]; exist{
        return []int{i,j}
    }
    m[n]=i
   }
   return nil
}