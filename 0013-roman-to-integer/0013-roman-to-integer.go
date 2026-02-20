func romanToInt(s string) int {
romans:=map[rune]int{
    'I':1,
    'V':5,
    'X':10,
    'L':50,
    'C':100,
    'D':500,
    'M':1000,
}
res:=0
for i:=0;i<len(s);i++{
    curr:=rune(s[i])
    if i+1 < len(s)  && romans[curr]<romans[rune(s[i+1])]{
        res-=romans[curr]
    }else{res+=romans[curr] }
}
return res
}