func hasDuplicate(nums []int) bool {
    m := make(map[int]int)

    for _,val := range nums{
        m[val]++
        if m[val] > 1{
            return true
        }
    }

    return false
}
