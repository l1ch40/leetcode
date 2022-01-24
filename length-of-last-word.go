func lengthOfLastWord(s string) int {
    vs := strings.TrimSpace(s)
    ss := strings.Split(vs, " ")
    return len(ss[len(ss) - 1])
}
