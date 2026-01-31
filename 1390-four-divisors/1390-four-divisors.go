func sumFourDivisors(nums []int) int {
    res := 0
    for _, n := range nums {
        sum := 0
        count := 0
        for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
            if n%i == 0 {
                count++
                sum += i
                if i != n/i {
                    count++
                    sum += n / i
                }
            }
        }
        if count == 4 {res += sum}
    }
    return res
}
