func findMaxConsecutiveOnes(nums []int) int {
    count,res:=0,0
    for _,n:=range nums{ 
        if n==0{
            count=0
        }else{count++}
    
     if count>res{res=count}
    }
    return res
}