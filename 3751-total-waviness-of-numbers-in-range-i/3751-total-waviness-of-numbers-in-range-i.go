func totalWaviness(num1 int, num2 int) int {
    count:=0
    if num1<100 && num2<100{
        return 0
    }
    for i:=num1;i<=num2;i++{
        nums:=[]int{}
        n:=i
        for n>0{
            nums=append(nums,n%10)
            n/=10
        }
        Reverse(&nums)

        for j:=1;j< len(nums)-1;j++{
            left:=nums[j-1]
            mid:=nums[j]
            right:=nums[j+1]
            if mid>left && mid>right{ count++ }
            if mid<left && mid<right{ count++ }
        }
    }
    return count
    
    }
    func Reverse(nums *[]int){
        for l, r := 0, len(*nums)-1; l < r; l, r = l+1, r-1 {
            (*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
        }
    }