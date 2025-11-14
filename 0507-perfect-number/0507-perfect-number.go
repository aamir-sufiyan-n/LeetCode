func checkPerfectNumber(num int) bool {
    sum:=0
    for i:=1;i<=num/2;i++{
        
        if num%i==0{sum+=i}
    }
    
    return num==sum
}