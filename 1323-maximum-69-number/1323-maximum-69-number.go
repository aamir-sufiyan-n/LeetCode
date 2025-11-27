func maximum69Number (num int) int {
    place:=1
    toAdd:=0
    temp:=num
    for temp>0{
        digit:=temp%10
        if digit==6 { toAdd=3*place }
        temp/=10
        place*=10
    }
    return num+toAdd
}