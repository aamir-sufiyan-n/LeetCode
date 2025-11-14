func checkPerfectNumber(num int) bool {
    sum:=1
    for i:=2;i*i<=num;i++{
        if num%i==0{
            sum+=i
            if i != num/i { sum+=num/i}
        }
    }
    if sum<=1 {return false}
    return num==sum
}