func arithmeticTriplets(nums []int, diff int) int {
checker:=make(map[int]bool)
count:=0
for _,n:=range nums{ checker[n]=true }
for _,n:=range nums{
  if checker[n+diff] && checker[n-diff]{
    count++
    }
}
return count
}