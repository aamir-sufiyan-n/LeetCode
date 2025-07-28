/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let boundary=0;
    let times=0;
    nums.forEach(element => {
     boundary+=element
     if(boundary==0)
        times+=1      
    });
     
     return times
};