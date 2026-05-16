func countNegatives(grid [][]int) int {
    // count:=0
    // for _,arr:=range grid{
    //     for _,n:=range arr{
    //         if n<0{count++}
    //     }
    // }
    // return count
    
    rows:=len(grid)
    cols:=len(grid[0])
    count:=0
    r,c:=0,cols-1
    for r<rows && c>=0{
        if grid[r][c]<0{
            count += rows-r
            c--
        }else{
            r++
        }
    }
    return count
}