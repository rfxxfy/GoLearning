func average(mass []int) float64 {
    sum := 0
    for _, m := range mass {
        sum += m
    }
    return float64(sum) / float64(len(mass))
}
