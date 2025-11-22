/**
 * @param {number[]} target
 * @param {number} n
 * @return {string[]}
 */
var buildArray = function(target, n) {
    let res=[]
    let j=0
    for (let i=1;i<=n&&j<target.length;i++){
        if (target.includes(i)){
            res.push("Push")
            j++
        }else{
            res.push("Push","Pop")
        }
    }
    return res
    
};