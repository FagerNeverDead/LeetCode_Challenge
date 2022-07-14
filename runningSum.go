func runningSum(nums []int) []int {
    count :=0
    arr := make([]int,len(nums))
    for i,v := range nums {
        count +=v
        arr[i] = count
    }
    return arr
}
