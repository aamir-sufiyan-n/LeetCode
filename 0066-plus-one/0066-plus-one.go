func plusOne(digits []int) []int {
    l:=len(digits)-1
    for i:=l;i>=0;i--{
        if digits[i]!=9{
        digits[i]+=1
        break }else{ digits[i]=0}
    }
    if digits[1]==0{
        digits=append([]int{1},digits...)
        return digits
    }else{return digits}
}