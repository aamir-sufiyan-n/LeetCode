func canConstruct(ransomNote string, magazine string) bool {
ranfreq:=make(map[rune]int)
for _,l:=range ransomNote{ ranfreq[l]++}
magfreq:=make(map[rune]int)
for _,l:=range magazine{ magfreq[l]++}

for _,c:= range ransomNote{
    if ranfreq[c] > magfreq[c] { return false}
}
return true
}