func totalFruit(fruits []int) int {
    bucketsNum := 3
    loadedFruite := make(map[int]int, 0)
    ans, i := 0, 0
    for inx, fruiteType := range fruits {
        loadedFruite[fruiteType] += 1
        for len(loadedFruite) >= bucketsNum {
            loadedFruite[fruits[i]] -= 1
            if loadedFruite[fruits[i]] == 0 {
                delete(loadedFruite, fruits[i])
            }
            i += 1
        }
        ans = max(ans, inx - i + 1)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
