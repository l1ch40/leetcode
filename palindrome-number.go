func isPalindrome(x int) bool {
    sx := strconv.Itoa(x)
    rsx := ReverseString(sx)
    if sx == rsx {
        return true
    }
    return false
}

func ReverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes) - 1; from < to; from, to = from + 1, to - 1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}
