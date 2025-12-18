func firstUniqChar(s string) int {
  look:=make(map[rune]int)
  for _,l:=range s{
    look[l]++
  }
  for i,l:=range s{
    if look[l]==1{
    return i
    }
  }
  return -1
}