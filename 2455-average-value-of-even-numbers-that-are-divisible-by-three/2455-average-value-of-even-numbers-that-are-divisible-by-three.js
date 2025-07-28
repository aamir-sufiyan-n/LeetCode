/**
 * @param {number[]} nums
 * @return {number}
 */
var averageValue = function(nums) {
      let filtered= nums.filter((Element)=>{
        return Element%2==0&& Element%3==0
    });
    let result=filtered.reduce((previous,current)=>{
        return previous+=current
    },0)
    if(result>0)
     result=result/filtered.length
   return Math.floor(result)
   
};