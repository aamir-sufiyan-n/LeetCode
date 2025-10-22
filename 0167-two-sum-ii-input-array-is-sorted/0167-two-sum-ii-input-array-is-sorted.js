/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    
    // for(let i=0;i<numbers.length;i++)
    // {
    //     for(let j=i+1;j<numbers.length;j++)
    //     if((numbers[i]+numbers[j])==target)
    //     return [i+1,j+1]
    // }
    let left=0;
    let right= numbers.length-1
    while (left<right)
    {
        let sum=numbers[left]+numbers[right]
        if(sum==target) 
        return [left+1,right+1]
        if(sum>target)right--
        else left++
    }
};