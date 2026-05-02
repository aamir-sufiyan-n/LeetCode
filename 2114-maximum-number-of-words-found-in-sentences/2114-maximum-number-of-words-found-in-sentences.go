func mostWordsFound(sentences []string) int {
    res:=0
    for _,s:=range sentences{
        count:=0
        words:=strings.Fields(s)
        for _,_=range words{
            count++
        }
        if count>res { res=count }
    }
    return res
}