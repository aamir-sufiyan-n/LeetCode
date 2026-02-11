func isPalindrome(s string) bool {
    new:=[]rune{}
    for _,c:=range []rune(s){
        if unicode.IsLetter(c) || unicode.IsNumber(c){
            new=append(new,unicode.ToLower(c))
        }
    }
    l,r:=0,len(new)-1
    for l<r{
        if new[l]!=new[r]{return false}
        l++
        r--
    }
    return true

}