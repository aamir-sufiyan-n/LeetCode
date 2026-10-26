func pivotInteger(n int) int {
   sum:=(n*(n+1))/2
   for i:=1;i<=n;i++{
     sum_left:=(i*(i+1))/2
     sum_right:= sum - (i*(i-1))/2
     if sum_left == sum_right{ return i }
   }
   return -1
}