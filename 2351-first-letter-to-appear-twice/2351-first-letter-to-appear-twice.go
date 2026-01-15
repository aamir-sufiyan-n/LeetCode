func repeatedCharacter(s string) byte {
    var checker=make(map[rune]int)
    for _,c:=range s{
        if _,ok:=checker[c]; ok{
            return byte(c)
        }
        checker[c]++
    }
    return 0
}