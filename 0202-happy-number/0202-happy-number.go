
func isHappy(n int) bool {
   seen:=make(map[int]bool)
   for n!=1 {
        if seen[n]{return false}
        seen[n]=true
        res:=0
        for n > 0{
            digit:=n%10
            res+=digit*digit
            n/=10
        }
        n=res
   }
   return true
}