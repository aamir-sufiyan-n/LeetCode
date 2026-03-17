func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    count:=0
    for i:=0;i<len(seats);i++{
        diff:=seats[i]-students[i]
        if diff < 0 {diff=-diff}
        count+=diff
    }
    return count
}