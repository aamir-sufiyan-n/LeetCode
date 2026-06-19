func largestAltitude(gain []int) int {
    alt:=0
    s:=0
    for _,n:=range gain{
      s+=n
      if s>alt{ alt = s}
    }
    return alt
}