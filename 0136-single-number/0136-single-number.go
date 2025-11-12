func singleNumber(nums []int) int {
    counter := make(map[int]int)
    for _,n:= range nums{
        counter[n]++
    }
    for num,count:= range counter{
        if count==1{
            return num
        }
    }
    return 0
}