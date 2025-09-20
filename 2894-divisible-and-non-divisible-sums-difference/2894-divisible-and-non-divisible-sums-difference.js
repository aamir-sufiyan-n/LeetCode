/**
 * @param {number} n
 * @param {number} m
 * @return {number}
 */
var differenceOfSums = function(n, m) {
  let divisible=0
  let nondivisible=0
  for(let i=0;i<=n;i++)
  {
    if(i%m==0)
    divisible+=i
    else nondivisible+=i
  }  
  return nondivisible-divisible
};