func addDigits(num int) int {
    if num<10 {return num}
    n:=num%10
    num/=10
    res:=num+n
    return addDigits(res)
}