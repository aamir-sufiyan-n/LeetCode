/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let position=0;
    let boundary=0;
    nums.forEach(element => {
     position+=element
     if(position==0)
        boundary+=1      
    });
     
     return boundary
};