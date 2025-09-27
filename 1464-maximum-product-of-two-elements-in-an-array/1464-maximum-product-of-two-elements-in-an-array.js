/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
    let sorted=nums.sort()
    let l=nums.length-1
    return ((sorted[l]-1)*(sorted[l-1]-1))
};