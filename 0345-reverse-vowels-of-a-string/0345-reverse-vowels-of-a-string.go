func reverseVowels(s string) string {
    wovels:=map[rune]bool{
        'a':true,'A':true,
        'e':true,'E':true,
        'i':true,'I':true,
        'o':true,'O':true,
        'u':true,'U':true,
    }
    r:=[]rune(s)
    i:=0
    j:=len(r)-1
    for i<=j{
        if wovels[r[i]] && wovels[r[j]]{
            r[i],r[j]=r[j],r[i]
            i++
            j--
        }else if wovels[r[i]] { j-- } else if wovels[r[j]] {
             i++ }else{
                i++
                j-- } 
    }
    return string(r)
}