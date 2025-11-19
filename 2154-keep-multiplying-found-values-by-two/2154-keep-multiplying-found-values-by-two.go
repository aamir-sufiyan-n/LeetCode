func findFinalValue(nums []int, original int) int {
    for _,n:= range nums{
        if n==original{
            original=n*2   
        }
    }
    return original
}