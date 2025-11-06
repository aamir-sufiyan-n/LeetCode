func maxFreqSum(s string) int {
    tracker:=make(map[rune]int)
    for _,ch:= range s { tracker[ch]++ }
    vowel:= map[rune]bool{
        'a':true, 'e':true ,'i':true, 'o':true, 'u':true,
        'A':true, 'E':true ,'I':true, 'O':true, 'U':true,
    }
    vowo:=0
    leto:=0
    for ch,freq:=range tracker{
        if vowel[ch]&&freq>vowo{
            vowo=freq
        }else if !vowel[ch] && freq>leto{leto=freq} 
    }
    return vowo+leto
}