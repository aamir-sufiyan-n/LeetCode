func isPalindrome(x int) bool {
    if x<0 ||(x%10==0 && x!=0){return false}
    reversed:=0
    for x > reversed{
        digit:=x%10
        reversed= reversed*10 +digit
        x/=10
    }
    return x==reversed || x==reversed/10
}