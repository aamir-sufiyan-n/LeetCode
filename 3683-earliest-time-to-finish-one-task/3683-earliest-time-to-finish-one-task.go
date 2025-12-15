func earliestTime(tasks [][]int) int {
    time:=math.MaxInt
    for _,task:= range tasks{
        time = min(time, task[0] + task[1])
    }
    return time
}