/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    let count=0
    let splitted=s.split("")
    for(let i=0;i<splitted.length;i++)
    {
        if(t.includes(splitted[i]))
        count++
    }
    if(count==s.length&&count==t.length)
    return true
    else return false
};