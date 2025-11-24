func pivotInteger(n int) int {
    sum:=(n * (n+1))/2
    for i:=1;i<=n;i++{
        suml:= (i * (i+1))/2
        sumr:= sum - (i*(i-1))/2
        if suml == sumr{
            return i
        }
    }
    return -1
}