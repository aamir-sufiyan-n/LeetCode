/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var minOperations = function(nums, k) {
    let count=0
    let a=nums.reduce((a,b)=>a+b,0)
    while(a%k!==0)
    {
        a--
        count++
    }
    return count
};