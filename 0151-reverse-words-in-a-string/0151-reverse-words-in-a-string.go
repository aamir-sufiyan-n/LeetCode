func reverseWords(s string) string {
    splitted:=strings.Fields(s) 
    return strings.Join(reverse(splitted)," ")
}
func reverse(s []string)[]string{
    i:=0
    j:=len(s)-1
    for i<j {
        s[i],s[j]=s[j],s[i]
        i++
        j--
    }
    return s
}