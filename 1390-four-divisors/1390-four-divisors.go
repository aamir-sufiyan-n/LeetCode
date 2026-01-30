func sumFourDivisors(nums []int) int {
    res:=0
    for _,n:=range nums{
        sum:=0
        count:=0
        for i:=1;i<=n;i++{
            if n%i==0{
                count++
                sum+=i
            }
        }
        if count==4{
            res+=sum
            sum=0
            count=0
        }
            sum=0
            count=0
    }
    return res
}