func addDigits(num int) int {
    if num<10 {return num}
    num1:=num%10
    num/=10
    res:=num1+num
    return addDigits(res)
}