func reversePrefix(s string, k int) string {
    i:=0
    j:=k-1
    c:=[]rune(s)
    for i<j{
        c[i],c[j]=c[j],c[i]
        i++
        j--
    }
    return string(c)
}