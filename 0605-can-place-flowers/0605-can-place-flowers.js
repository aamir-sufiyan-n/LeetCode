/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @return {boolean}
 */
var canPlaceFlowers = function(flowerbed, n) {
    let beds=[0,...flowerbed,0]
  for(let i=1;i<beds.length-1;i++)
   {
    if(beds[i]+beds[i+1]+beds[i-1]==0)
    {
       n-- 
       beds[i]=1    
       }
    }
     return n<=0
};