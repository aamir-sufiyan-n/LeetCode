/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    let count=0;
    let prev=0;
    let majority=0;
    let n=nums.length/2
    for(let i=0;i<nums.length;i++)
    {
        for(let j=i;j<nums.length;j++)
        {
            if(nums[i]==nums[j])
            count++
        }

        if(count>n&&count>prev)
        {
            majority=nums[i]
            prev=count
        }
            count=0
    }
    return majority
};