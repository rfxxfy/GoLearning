func binarySearch(mass []int, x int) bool {
    l := 0
    r := len(mass)

    for l+1 < r {
        m := (l + r) / 2
        if mass[m] <= x {
            l = m
        } else {
            r = m
        }
    }

    return mass[l] == x
}
