/**
 * @param {number} x
 * @return {number}
 */
var sumOfTheDigitsOfHarshadNumber = function(x) {
    let sum=String(x).split("").map(Number).reduce((a,b)=>a+b,0)
    if(x%sum==0)
    return sum
    else
    return(-1)
};