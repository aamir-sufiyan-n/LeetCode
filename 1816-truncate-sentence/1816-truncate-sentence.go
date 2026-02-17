func truncateSentence(s string, k int) string {
    some:=strings.Fields(s)
    some=some[:k]
    return strings.Join(some," ")
}