func mergeAlternately(word1 string, word2 string) string {
    i:=0
    var res=[]rune{}
    w1:=[]rune(word1)
    w2:=[]rune(word2)
    for i < min(len(w1),len(w2)){
        res=append(res,w1[i])
        res=append(res,w2[i])
        i++
    } 
    res=append(res,w1[i:]...)
    res=append(res,w2[i:]...)
    return string(res)
}