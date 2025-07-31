/**
 * @param {number[]} nums
 * @return {number}
 */
var findGCD = function(nums) {
     let sorted=nums.sort((a,b)=>a-b)
    let mn=sorted[0]
    let mx=sorted[sorted.length-1]
    let count=[];
   for(let i=1;i<=mn;i++)
   {
    if(mn%i==0 && mx%i==0 )
          count.push(i)
   }
   return Math.max(...count)
   
};