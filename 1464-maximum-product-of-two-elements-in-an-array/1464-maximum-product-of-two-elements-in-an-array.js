/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
    let sorted=nums.sort((a,b)=>a-b)
    console.log(sorted)
    let l=nums.length-1
    return ((sorted[l]-1)*(sorted[l-1]-1))
};