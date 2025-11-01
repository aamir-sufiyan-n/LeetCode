func getSneakyNumbers(nums []int) []int {
   result:=[]int{}
   nos:=make(map[int]int)
   for _,v:= range nums{
    nos[v]++
   }
   for num,count:=range nos{
    if count>1{
        result=append(result,num)
    }
   }
   return result
}