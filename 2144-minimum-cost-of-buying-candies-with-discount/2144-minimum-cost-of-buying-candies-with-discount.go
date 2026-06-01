func minimumCost(cost []int) int {
    sort.Ints(cost)
    sum:=0
    for i:=len(cost)-1;i>=0;i--{
        if (len(cost)-i) %3 != 0{
            sum+=cost[i]
        }
    }
    return sum
}