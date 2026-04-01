func decodeMessage(key string, message string) string {
    l:='a'
    code:=make(map[rune]rune,26)
    code[' ']=' '
    for _,c:=range key{
        if _,ok:=code[c];!ok{
            code[c]=l
            l++
        }
    }
    res:=[]rune{}
    for _,c:=range message{
        res=append(res,code[c])
    }
    return string(res)
}