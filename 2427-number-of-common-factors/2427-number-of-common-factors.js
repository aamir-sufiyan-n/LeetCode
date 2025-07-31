/**
 * @param {number} a
 * @param {number} b
 * @return {number}
 */
var commonFactors = function(a, b) {
    let factors=0;
    for(let i=1;i<10000;i++)
    {
        if(b%i==0&&a%i==0)
            factors++
    }
    return factors
};