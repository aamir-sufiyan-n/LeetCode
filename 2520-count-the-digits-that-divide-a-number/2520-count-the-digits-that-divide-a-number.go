
func countDigits(num int) int {
    count := 0
    n := num
    for n > 0 {
        dig := n % 10
        if dig != 0 && num%dig == 0 {
            count++
        }
        n /= 10
    }
    return count
}
