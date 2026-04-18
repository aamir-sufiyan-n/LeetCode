func mirrorDistance(n int) int {
    num:=n
    rev:=0
    for n !=0{
        digit:= n%10
        rev = rev*10 + digit
        n/=10
    }
    return abs(num-rev)
}
func abs(n int)int{
    if n<0 { return n * -1}
    return n
}