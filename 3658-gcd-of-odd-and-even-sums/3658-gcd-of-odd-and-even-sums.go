func gcdOfOddEvenSums(n int) int {
    oddSum:=0
    evenSum:=0
    for i:=1 ; i<=n ; i++{
            evenSum += i * 2
            oddSum += (i*2)-1
    }
    return gcd(oddSum,evenSum)
}
func gcd(a,b int )int{
    for b!=0{
        a,b = b,a%b
    }
    return a
}