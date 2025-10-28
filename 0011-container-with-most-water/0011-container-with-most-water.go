func maxArea(height []int) int {
    left:=0;
    right:=len(height)-1
    var result int
    for left<=right{
        area:=(right-left)*min(height[left],height[right])
        if(height[left]<=height[right]){
            left++
        }else{
            right--
        }
        if(area>result){
        result=area
        }
    }
    return result
}