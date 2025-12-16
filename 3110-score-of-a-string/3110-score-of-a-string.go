func scoreOfString(s string) int {
    res:=0
    ascii:=[]byte(s)
    for i:=0;i<len(ascii)-1;i++{
        val:=int(ascii[i])-int(ascii[i+1])
        res+=abs(val)
    }
    return res
}

func abs(x int)int{
    if x > 0{return x}
    if x<0 {return x * -1}
    return x
}