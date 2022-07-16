func decompressRLElist(nums []int) []int {
    result := make([]int,lenD(nums))
    point := 0
    for l:=0;l<len(nums);l+=2{
        f := nums[l]
        v := nums[l+1]
        for f > 0 {
            result[point] = v
            f--
            point++
        }
    }
    return result
}
func lenD(nums []int) (lenArr int) {
    for l:=0;l<len(nums);l+=2{
        lenArr +=nums[l] 
    }
    return
}
