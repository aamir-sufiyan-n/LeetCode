func findTheDifference(s string, t string) byte {
   checker:=make(map[rune]int)
   for _,char:= range s{
    checker[char]++
   }
   for _,l:=range t{
    checker[l]--
   }
   for k,v:=range checker{
    if v<0{
        return byte(k)
    }
   }
   return byte(0)
}