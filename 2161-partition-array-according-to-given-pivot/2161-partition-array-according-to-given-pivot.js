/**
 * @param {number[]} nums
 * @param {number} pivot
 * @return {number[]}
 */
var pivotArray = function(nums, pivot) {
    let lessthan=[]
    let greater=[]
    let equal=[]
    for(let i=0;i<nums.length;i++)
    {
        if(nums[i]>pivot)
        greater.push(nums[i])
        else if(nums[i]==pivot)
        equal.push(nums[i])
        else 
        lessthan.push(nums[i])
    }
    let result=[...lessthan,...equal,...greater]
    return result
};