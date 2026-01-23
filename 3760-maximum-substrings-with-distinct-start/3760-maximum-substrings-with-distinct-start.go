func maxDistinct(s string) int {
    checker:=make(map[rune]int)
    for _,c:=range s{
        checker[c]++
    }
    return len(checker)
}