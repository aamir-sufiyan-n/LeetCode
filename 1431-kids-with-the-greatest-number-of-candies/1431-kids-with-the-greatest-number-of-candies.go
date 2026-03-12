func kidsWithCandies(candies []int, extraCandies int) []bool {
    var res[]bool
    maxVal := candies[0]
    for _,n:=range candies{
        if n>maxVal{ maxVal=n}
    }
    for _,n:=range candies{
        if n + extraCandies >= maxVal{
            res=append(res,true)
        }else{res=append(res,false)}
    }
    return res
}