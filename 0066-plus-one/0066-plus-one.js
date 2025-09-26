/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
  let l=digits.length-1
  for(let i=l;i>=0;i--)
  {
    if(digits[i]!=9)
    {
        digits[i]+=1
        break
    }
    else
        digits[i]=0
  }
if(digits[0]==0)
    return [1,...digits]
    else return digits
};
