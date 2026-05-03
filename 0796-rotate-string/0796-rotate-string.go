func rotateString(s string, goal string) bool {
    joined:= s + s
    if len(s) != len(goal){ return false }
    return strings.Contains(joined,goal)
}