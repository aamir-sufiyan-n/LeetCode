func maxDistance(colors []int) int {
    n:=len(colors)
    left:=0
    for i:=0;i < n;i++{
        if colors[i]!=colors[n-1]{
            left= (n-1) - i 
            break
        } 
    }
    right:=0
    for j:=n-1;j>=0;j--{
        if colors[j]!=colors[0]{
            right = j
            break
        }
    }
    return max(left,right)
}
