func mostWordsFound(sentences []string) int {
    res:=0
    for _,s:=range sentences{
        count:=0
        sen:=strings.Fields(s)
        for _,_=range sen{
            count++
        }
        if count>res{res=count}
    }
    return res
}