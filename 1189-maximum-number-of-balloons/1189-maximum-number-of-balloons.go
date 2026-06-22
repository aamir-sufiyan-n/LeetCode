func maxNumberOfBalloons(text string) int {
    track:=make(map[rune]int)
    minCount:=math.MaxInt32
    need:= map[rune]int{
        'b':1,'a':1,'l':2,'o':2,'n':1,
    }
    for _,c:= range text{
        track[c]++
    }
    for letter,req:=range need{
        avl:=track[letter]
        possible:=avl/req
        if possible<minCount{
            minCount=possible
        }
    }
    return minCount
}