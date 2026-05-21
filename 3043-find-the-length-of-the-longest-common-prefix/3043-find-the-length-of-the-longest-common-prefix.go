func longestCommonPrefix(arr1 []int, arr2 []int) int {
    seen:=make(map[int]bool)
    for _,n:=range arr1{
        x:=n
        for x>0{
            seen[x]=true
            x/=10
        }
    }
    maxLen:=0
    for _,num:=range arr2{
        x:=num
        for x>0{
            if seen[x]{
                len:=len(fmt.Sprint(x))
                if len>maxLen{maxLen=len}
            }
            x/=10
        }
    }
    return maxLen
}