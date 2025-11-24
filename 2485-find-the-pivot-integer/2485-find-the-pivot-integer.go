func pivotInteger(n int) int {
    // sum:=(n * (n+1))/2
    // for i:=1;i<=n;i++{
    //     suml:= (i * (i+1))/2
    //     sumr:= sum - (i*(i-1))/2
    //     if suml == sumr{
    //         return i
    //     }
    // }
    // return -1

    for i:=1;i<=n;i++{
        sum:=0
        sum2:=0
        for j:=1;j<=i;j++{
            sum+=j
        }
        for k:=i;k<=n;k++{
            sum2+=k
        }
        if sum==sum2{
            return i
        }else{
            sum=0
            sum2=0
        }
    }
    return -1
}