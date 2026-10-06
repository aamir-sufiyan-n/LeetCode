func minAddToMakeValid(s string) int {
    open, a := 0, 0
    for _, c := range s {
        if c == '(' {
            open++
        } else {
            if open > 0 { open-- } else { a++ }
        }
    }
    return a + open
}