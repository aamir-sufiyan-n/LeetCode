func findMaxConsecutiveOnes(nums []int) int {
    count,res:=0,0
    for _,n:=range nums{
        if n==1 {
            count++
            if count>res { res=count }
        }else{ count=0 } 
    }
    return res
}