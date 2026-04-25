func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    res:=0
    for _,n:=range hours{
        if n>=target{
            res++
        }
    }
    return res
}