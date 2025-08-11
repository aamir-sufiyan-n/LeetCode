/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
    let splitted=String(n).split("").map(Number)
    let sum=splitted.reduce((a,b)=>a+b,0)
    let product=splitted.reduce((a,b)=>a*b,1)
    let result=product-sum;
    return result
};