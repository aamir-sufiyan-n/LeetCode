func rotate(nums []int, k int) []int {
    k=k%len(nums)
    copy(nums,append(nums[len(nums)-k:],nums[:len(nums)-k]...))
    return nums
}