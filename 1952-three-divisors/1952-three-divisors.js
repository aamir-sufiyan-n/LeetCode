/**
 * @param {number} n
 * @return {boolean}
 */
var isThree = function(n) {
     let divisors=0
    for(let i=1;i<=n;i++){
        if(n%i===0)
            divisors+=1;
    }
    if(divisors===3)
        return true
    else 
        return false
};