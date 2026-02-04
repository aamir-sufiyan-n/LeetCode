func reverseVowels(s string) string {
    r:=[]rune(s)
    wovels:=[]rune{'a','A','e','E','i','I','o','O','u','U'}
    i:=0
    j:=len(r)-1
    for i<=j{
        if slices.Contains(wovels,r[i])&&slices.Contains(wovels,r[j]){
            r[i],r[j]=r[j],r[i]
            i++
            j--
        }else if slices.Contains(wovels,r[i]){
            j--}else if slices.Contains(wovels,r[j]){
                i++
            }else{
                i++
                j--
            } 
    }
    return string(r)
}