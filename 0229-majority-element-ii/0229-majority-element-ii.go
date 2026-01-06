func majorityElement(nums []int) []int {
    var res[] int
    checker:=make(map[int]int)
    for _,n:=range nums{
        checker[n]++
    }
    for num,freq:=range checker{
        if freq> (len(nums)/3){
            res=append(res,num)
        }
    }
    return res
}