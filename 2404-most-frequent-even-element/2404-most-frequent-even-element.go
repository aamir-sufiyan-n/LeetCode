func mostFrequentEven(nums []int) int {
 bestNum:=-1
 max:=0
 checker:=make(map[int]int)
 for _,n:=range nums{
    if n%2==0{
        checker[n]++
    }
 }
 for num,freq:=range checker{
    if freq>max{
        max=freq
        bestNum=num
    }
    if freq==max && num<bestNum{
        bestNum=num
    }
 }
 return bestNum
}