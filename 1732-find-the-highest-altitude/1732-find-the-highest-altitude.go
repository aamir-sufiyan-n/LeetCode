func largestAltitude(gain []int) int {
    alt,s:=0,0
    for _,n:=range gain{
      s+=n
      if s>alt{ alt = s}
    }
    return alt
}