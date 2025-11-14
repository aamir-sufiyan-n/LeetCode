func checkPerfectNumber(num int) bool {
    sum:=1
    for i:=2;i<=num/2;i++{
        
        if num%i==0{sum+=i}
    }
    if sum==1{return false}
    return num==sum
}