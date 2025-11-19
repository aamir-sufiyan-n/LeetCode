func findFinalValue(nums []int, original int) int {
 track:=make(map[int]bool)
 for _,n:=range nums{
    track[n]=true
 }
 for track[original] {
     original*=2 
     }
     
return original
}