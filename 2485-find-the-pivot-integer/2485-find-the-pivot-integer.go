func pivotInteger(n int) int {
    sum1:=0
    sum2:=0
    for i:=1;i<=n;i++{
        sum1+=i
        for j:=i;j<=n;j++{
            sum2+=j
        }
        if sum1==sum2{return i}
        sum2=0
    }
   return -1
}