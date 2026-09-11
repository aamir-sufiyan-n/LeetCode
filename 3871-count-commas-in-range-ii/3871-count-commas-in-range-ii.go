func countCommas(n int64) int64 {
    var ans int64
    for _,x:= range []int64{
        999,
        999999,
        999999999,
        999999999999,
        999999999999999,
    }{
        if n>x { ans += n-x }
    }
    return ans
}