/**
 * @param {number[]} nums
 * @return {number}
 */
var minElement = function(nums) {
    let result=[]
    let sum=0
    for(let i=0;i<nums.length;i++)
    {
        let num=String(nums[i]).split("")
        for(let j=0;j<num.length;j++)
        {
            sum+=Number(num[j])
        }
        result.push(sum)
        sum=0
    }
    return Number(Math.min(...result))
};