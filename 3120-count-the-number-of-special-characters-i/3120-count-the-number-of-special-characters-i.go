func numberOfSpecialChars(word string) int {
    check:=make(map[rune]bool)
     for _,c:=range word{
          check[c]=true
    }
    count:=0
    for i:='a';i<='z';i++{
        if check[i] && check[unicode.ToUpper(i)]{
            count++
        }
    }
    return count 
}