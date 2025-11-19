func checkIfExist(arr []int) bool {
    track:=make(map[int]bool)
    for _,n:=range arr{
        if track[n*2] || (n%2==0 && track[n/2] ){
            return true
        }
        track[n]=true
    }
    return false
}