func maxIceCream(costs []int, coins int) int {
    count:=0
    sort.Ints(costs)
    for _,p:=range costs{
        if coins-p >= 0{
            coins-=p
            count++
        }
    }
    return count
}