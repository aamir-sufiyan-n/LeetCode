func numberOfSpecialChars(word string) int {
    // upper:=make(map[rune]bool)
    // lower:=make(map[rune]bool)
    // for _,c:=range word{
    //     if unicode.IsUpper(c){
    //       upper[c]=true
    //     }
    //     if unicode.IsLower(c){
    //       lower[c]=true
    //     }
    // }
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