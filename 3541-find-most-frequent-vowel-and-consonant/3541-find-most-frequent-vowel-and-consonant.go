func maxFreqSum(s string) int {
    tracker:=make(map[rune]int)
    for _,ch:= range s { tracker[ch]++ }
    vowo:=0
    leto:=0
    for ch,freq:=range tracker{
        if strings.ContainsRune("aeiou",ch){
            if   freq>vowo { vowo=freq }
        }else {
            if freq>leto { leto=freq }  
        }
    }
    return vowo+leto
}